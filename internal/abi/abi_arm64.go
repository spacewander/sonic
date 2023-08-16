/*
 * Copyright 2022 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package abi

import (
    `fmt`
    `reflect`

    . `github.com/chenzhuoyu/iasm/arch/aarch64`
    `github.com/chenzhuoyu/iasm/asm`
)

const (
    PtrSize  = 8    // pointer size
    PtrAlign = 8    // pointer alignment
)

func (self *Frame) Offs() uint32 {
    return PtrSize + uint32(len(ReservedRegs(self.ccall)) * PtrSize + len(self.locals)*PtrSize)
}

func (self *Frame) argv(i int) *asm.MemoryOperand {
    return Ptr(SP, int32(self.Prev() + self.desc.Args[i].Mem))
}

// spillv is used for growstack spill registers
func (self *Frame) spillv(i int) *asm.MemoryOperand {
    // remain one slot for caller return pc
    return Ptr(SP, PtrSize + int32(self.desc.Args[i].Mem))
}

func (self *Frame) retv(i int) *asm.MemoryOperand {
    return Ptr(SP, int32(self.Prev() + self.desc.Rets[i].Mem))
}

func (self *Frame) resv(i int) *asm.MemoryOperand {
    return Ptr(SP, int32(self.Offs() - uint32((i+1) * PtrSize)))
}

func (self *Frame) emitGrowStack(p *Program, entry *asm.Label) {
    // spill all register arguments
    for i, v := range self.desc.Args {
        if v.InRegister {
            p.STR(v.Reg, self.spillv(i))
        }
    }

    // call runtime.morestack_noctxt
    loadBigImm(p, uint64(F_morestack_noctxt), X12)
    p.BLR(X12)

    // load all register arguments
    for i, v := range self.desc.Args {
        if v.InRegister {
            p.LDR(v.Reg, self.spillv(i))
        }
    }

    // jump back to the function entry
    p.BL(entry)
}

func (self *Frame) GrowStackTextSize() uint32 {
    p := Builder(asm.GetArch("aarch64").CreateProgram())
    l := asm.CreateLabel("entry")
    self.emitGrowStack(p, l)
    p.Link(l)
    return uint32(len(p.Assemble(0)))
}


func (self *Frame) emitPrologue(p *Program) {
    p.STR(LR, Ptr(SP, -int(self.Size()), PostIndex)) // str    x30, [sp, #-{size}]!
    p.STUR(FP, Ptr(SP, -8))                     // stur   x29, [sp, #-0x8]
    p.SUB(FP, SP, 8)                            // sub    x29, sp, #0x8
}

func (self *Frame) emitEpilogue(p *Program) {
    p.LDP(FP, LR, Ptr(SP, -8)) // ldp    x29, x30, [sp, #-0x8]
    p.ADD(SP, SP, self.Size()) // add    sp, sp, #{size}
    p.RET()                    // ret
}

func (self *Frame) emitReserveRegs(p *Program) {
    // spill reserved registers
    for i, r := range ReservedRegs(self.ccall) {
        switch r.(type) {
        case XRegister, WRegister:
            p.STR(r, self.resv(i))
        default:
            panic(fmt.Sprintf("unsupported register type %t to reserve", r))
        }
    }
}

func (self *Frame) emitSpillPtrs(p *Program) {
    // spill pointer argument registers
    for i, r := range self.desc.Args {
        if r.InRegister && r.IsPointer {
            p.STR(r.Reg, self.argv(i))
        }
    }
}

func (self *Frame) emitClearPtrs(p *Program) {
    // spill pointer argument registers
    for i, r := range self.desc.Args {
        if r.InRegister && r.IsPointer {
            p.STR(XZR, self.argv(i))
        }
    }
}

// addr must be the pointer to store PC
func (self *Frame) emitCallC(p *Program, pc uintptr) {
    loadBigImm(p, uint64(pc), X12)
    p.BLR(X12)
}

func loadBigImm(p *Program, imm uint64, reg interface{}) {
    p.MOVZ(reg, uint16(imm|0xffff), LSL(0))
    p.MOVK(reg, uint16((imm>>16)|0xffff), LSL(16))
    p.MOVK(reg, uint16((imm>>32)|0xffff), LSL(32))
    p.MOVK(reg, uint16((imm>>48)|0xffff), LSL(48))
}

type floatKind uint8

const (
    notFloatKind floatKind = iota
    floatKind32
    floatKind64
)

type Parameter struct {
    InRegister bool
    IsPointer  bool
    IsFloat    floatKind
    Reg        interface{}
    Mem        uint32
    Type       reflect.Type
}

func mkIReg(vt reflect.Type, reg XRegister) (p Parameter) {
    p.Reg = reg
    p.Type = vt
    p.InRegister = true
    p.IsPointer = isPointer(vt)
    return
}

func isFloat(vt reflect.Type) floatKind {
    switch vt.Kind() {
    case reflect.Float32:
        return floatKind32
    case reflect.Float64:
        return floatKind64
    default:
        return notFloatKind
    }
}

func mkXReg(vt reflect.Type, i int) (p Parameter) {
    p.Type = vt
    p.InRegister = true
    p.IsFloat = isFloat(vt)
    if p.IsFloat == floatKind32 {
        p.Reg = xregOrderGo[i]
    } else if p.IsFloat == floatKind64 {
        p.Reg = yregOrderGo[i]
    } else {
        panic("not a float type!")
    }
    return
}

func mkStack(vt reflect.Type, mem uint32) (p Parameter) {
    p.Mem = mem
    p.Type = vt
    p.InRegister = false
    p.IsPointer = isPointer(vt)
    p.IsFloat = isFloat(vt)
    return
}

func (self Parameter) String() string {
    if self.InRegister {
        return fmt.Sprintf("[%%%s, Pointer(%v), Float(%v)]", self.Reg, self.IsPointer, self.IsFloat)
    } else {
        return fmt.Sprintf("[%d(FP), Pointer(%v), Float(%v)]", self.Mem, self.IsPointer, self.IsFloat)
    }
}

func CallC(pc uintptr, fr Frame, maxStack uintptr) []byte {
    p := Builder(asm.GetArch("aarch64").CreateProgram())
    stack := asm.CreateLabel("_stack_grow")
    entry := asm.CreateLabel("_entry")
    p.Link(entry)
    fr.emitStackCheck(p, stack, maxStack)
    fr.emitPrologue(p)
    fr.emitReserveRegs(p)
    fr.emitSpillPtrs(p)
    fr.emitExchangeArgs(p)
    fr.emitCallC(p, pc)
    fr.emitRestoreRegs(p)
    fr.emitEpilogue(p)
    p.Link(stack)
    fr.emitGrowStack(p, entry)

    return p.Assemble(0)
}


func (self *Frame) emitDebug(p *Program) {
    p.BRK(0)
}