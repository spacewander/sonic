//go:build go1.17
// +build go1.17

/*
 * Copyright 2021 ByteDance Inc.
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

package loader

import (
	"runtime"
	"runtime/debug"
	"strconv"
	"testing"
	"unsafe"

	"github.com/bytedance/sonic/internal/rt"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
    defer func() {
        if r := recover(); r != nil {
            runtime.GC()
            if r != "hook1" {
                t.Fatal("not right panic:" + r.(string))
            }
        } else {
            t.Fatal("not panic")
        }
    }()
    var hstr string

    type TestFunc func(i *int, hook func(i *int)) int
    var hook = func(i *int) {
        runtime.GC()
        debug.FreeOSMemory()
        hstr = ("hook" + strconv.Itoa(*i))
        runtime.GC()
        debug.FreeOSMemory()
        panic(hstr)
    }
    // var f TestFunc = func(i *int, hook func(i *int)) int {
    //     var t = *i
    //     hook(i)
    //     return t + *i
    // }
    //
    text := []byte{
        0xfe, 0x0f, 0x1e, 0xf8,    // str x30, [sp, #-32]!
        0xfd, 0x83, 0x1f, 0xf8,    // str x29, [sp, #-8]
        0xfd, 0x23, 0x00, 0xd1,    // sub x29, sp, #8
        0xe0, 0x07, 0x00, 0xf9,    // str x0, [sp, #8]  # i
        0x02, 0x00, 0x40, 0xf9,    // ldr x2, [x0]
        0xe2, 0x0b, 0x00, 0xf9,    // str x2, [sp, #16] # t
        0xfa, 0x03, 0x01, 0xaa,    // mov x26, x1       # closure pointer
        0x23, 0x00, 0x40, 0xf9,    // ldr x3, [x1]
        0x60, 0x00, 0x3f, 0xd6,    // blr  x3
        0xe1, 0x07, 0x40, 0xf9,    // ldr x1, [sp, #8]
        0x21, 0x00, 0x40, 0xf9,    // ldr x1, [x1]
        0xe2, 0x0b, 0x40, 0xf9,    // ldr x2, [sp, #16]
        0x20, 0x00, 0x02, 0x8b,    // add x0, x1, x2
        0xfd, 0xfb, 0x7f, 0xa9,    // ldp x29, x30, [sp, #-8]
        0xff, 0x83, 0x00, 0x91,    // add sp, sp, #32
        0xc0, 0x03, 0x5f, 0xd6,    // ret
    }
    // text, err := hex.DecodeString(src)
    // if err != nil {
    //     t.Fatal(err)
    // }
    size := uint32(len(text))
    fn := Func{
        ID: 0,
        Flag: 0,
        ArgsSize: 16,
        EntryOff: 0,
        TextSize: size,
        DeferReturn: 0,
        FileIndex: 0,
        Name: "dummy",
    }

    fn.Pcsp = &Pcdata{
        {PC: size-4, Val: 32},
        {PC: size, Val: 0},
    }

    fn.Pcline = &Pcdata{
        {PC: 0x20, Val: 1},
        {PC: 0x40, Val: 2},
        {PC: size, Val: 3},
    }

    fn.Pcfile = &Pcdata{
        {PC: size, Val: 0},
    }

    fn.PcUnsafePoint = &Pcdata{
        {PC: size, Val: PCDATA_UnsafePointUnsafe},
    }

    fn.PcStackMapIndex = &Pcdata{
        {PC: size, Val: 0},
    }

    args := rt.StackMapBuilder{}
    args.AddField(true)
    args.AddField(true)
    fn.ArgsPointerMaps = args.Build()

    locals := rt.StackMapBuilder{}
    locals.AddField(false)
    locals.AddField(false)
    fn.LocalsPointerMaps = locals.Build()

    rets := Load(text, []Func{fn}, "dummy_module", []string{"github.com/bytedance/sonic/dummy.go"})
    println("func address ", *(*unsafe.Pointer)(rets[0]))

    testFunc(func(){
        f := *(*TestFunc)(unsafe.Pointer(&rets[0]))
        i := 1
        j := f(&i, hook)
        require.Equal(t, 2, j)
        require.Equal(t, "hook1", hstr)
    })
   
    fi := runtime.FuncForPC(*(*uintptr)(rets[0]))
    require.Equal(t, "dummy", fi.Name())
    file, line := fi.FileLine(0)
    require.Equal(t, "github.com/bytedance/sonic/dummy.go", file)
    require.Equal(t, 0, line)
}

//go:noinline
func testFunc( f func()) {
    f()
}