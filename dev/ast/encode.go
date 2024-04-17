
package ast

import (
    `runtime`
    `unsafe`
    `sync`

    `github.com/bytedance/sonic/internal/native`
    `github.com/bytedance/sonic/internal/rt`
)

const (
    _MaxBuffer = 1024    // 1KB buffer size
)


var typeByte = rt.UnpackEface(byte(0)).Type
var bytesPool = sync.Pool{}

func (self *Node) MarshalJSON() ([]byte, error) {
    buf := newBuffer()
    err := self.encode(buf)
    if err != nil {
        freeBuffer(buf)
        return nil, err
    }

    ret := make([]byte, len(*buf))
    copy(ret, *buf)
    freeBuffer(buf)
    return ret, err
}

func newBuffer() *[]byte {
    if ret := bytesPool.Get(); ret != nil {
        return ret.(*[]byte)
    } else {
        buf := make([]byte, 0, _MaxBuffer)
        return &buf
    }
}

func freeBuffer(buf *[]byte) {
    *buf = (*buf)[:0]
    bytesPool.Put(buf)
}

func (self *Node) encode(buf *[]byte) error {
    if self.IsRaw() {
        return self.encodeRaw(buf)
    }

    switch self.getType() {
		case _V_NONE  : return ErrNotExist
        case _V_NULL  : return self.encodeNull(buf)
        case _V_ERROR : return self.getError()
        case _V_TRUE  : return self.encodeTrue(buf)
        case _V_FALSE : return self.encodeFalse(buf)

		case _V_INT64 : return self.encodeInt(buf)
		case _V_UINT64: return self.encodeUint(buf)
		case _V_FLOAT64: return self.encodeUint(buf)
        case _V_NUMBER: return self.encodeNumber(buf)

        case _V_ARRAY : return self.encodeArray(buf)
        case _V_OBJECT: return self.encodeObject(buf)
        case _V_STRING: return self.encodeString(buf)
        case _V_ANY   : panic("to do")
	
        default      : return ErrUnsupportType 
    }
}

func (self *Node) encodeRaw(buf *[]byte) error {
    *buf = append(*buf, self.getRaw()...)
    return nil
}

func (self *Node) encodeNull(buf *[]byte) error {
    *buf = append(*buf, bytesNull...)
    return nil
}

func (self *Node) encodeInt(buf *[]byte) error {
	panic("todo")
}

func (self *Node) encodeUint(buf *[]byte) error {
	panic("todo")
}

func (self *Node) encodeFloat(buf *[]byte) error {
	panic("todo")
 }

func (self *Node) encodeTrue(buf *[]byte) error {
    *buf = append(*buf, bytesTrue...)
    return nil
}

func (self *Node) encodeFalse(buf *[]byte) error {
    *buf = append(*buf, bytesFalse...)
    return nil
}

func (self *Node) encodeNumber(buf *[]byte) error {
    *buf = append(*buf, self.getNumber()...)
    return nil
}

func (self *Node) encodeString(buf *[]byte) error {
    quote(buf, self.getString())
    return nil
}

func (self *Node) encodeArray(buf *[]byte) error {
    nb := self.getLen()
    if nb == 0 {
        *buf = append(*buf, bytesArray...)
        return nil
    }
    
	println("nb is ", nb)

    *buf = append(*buf, '[')

    var started bool
    for i := 0; i < nb; i++ {
        n := self.getElem(i)
        if !n.Exists() {
            continue
        }
        if started {
            *buf = append(*buf, ',')
        }
        started = true
        if err := n.encode(buf); err != nil {
            return err
        }
    }

    *buf = append(*buf, ']')
    return nil
}

func (self *Pair) encode(buf *[]byte) error {
    if len(*buf) == 0 {
        *buf = append(*buf, '"', '"', ':')
        return self.Value.encode(buf)
    }

    quote(buf, self.Key)
    *buf = append(*buf, ':')

    return self.Value.encode(buf)
}

func (self *Node) encodeObject(buf *[]byte) error {
    nb := self.getLen()
    if nb == 0 {
        *buf = append(*buf, bytesEmptyObject...)
        return nil
    }
    
    *buf = append(*buf, '{')

    var started bool
    for i := 0; i < nb; i++ {
        n := self.getPair(i)
        if n == nil || !n.Value.Exists() {
            continue
        }
	
        if started {
            *buf = append(*buf, ',')
        }
	
        started = true
        if err := n.encode(buf); err != nil {
            return err
        }
    }

    *buf = append(*buf, '}')
    return nil
}

//go:nocheckptr
func quote(buf *[]byte, val string) {
    *buf = append(*buf, '"')
    if len(val) == 0 {
        *buf = append(*buf, '"')
        return
    }

    sp := rt.IndexChar(val, 0)
    nb := len(val)
    b := (*rt.GoSlice)(unsafe.Pointer(buf))

    // input buffer
    for nb > 0 {
        // output buffer
        dp := unsafe.Pointer(uintptr(b.Ptr) + uintptr(b.Len))
        dn := b.Cap - b.Len
        // call native.Quote, dn is byte count it outputs
        ret := native.Quote(sp, nb, dp, &dn, 0)
        // update *buf length
        b.Len += dn

        // no need more output
        if ret >= 0 {
            break
        }

        // double buf size
        *b = growslice(typeByte, *b, b.Cap*2)
        // ret is the complement of consumed input
        ret = ^ret
        // update input buffer
        nb -= ret
        sp = unsafe.Pointer(uintptr(sp) + uintptr(ret))
    }

    runtime.KeepAlive(buf)
    runtime.KeepAlive(sp)
    *buf = append(*buf, '"')
}