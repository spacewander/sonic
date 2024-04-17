package ast

import (
	"unsafe"
	"encoding/json"

	"github.com/bytedance/sonic/internal/rt"
)


/***************** Get Helper ***********************/

func (self *Node) getType() valueType {
	return self.typ & _TYPE_MASK
}

func (self *Node) getLen() int {
	return int(self.val)
}

func (self *Node) isRaw() bool {
	return (self.typ&_RAW_MASK) != 0
}

func (self *Node) getRaw() string {
	return rt.StrFrom(self.ptr, int64(self.getLen()))
}

func (self *Node) getError() error {
	return self
}

func (self *Node) getString() string {
	return rt.StrFrom(self.ptr, int64(self.getLen()))
}

func (self *Node) getNumber() json.Number {
	return json.Number(rt.StrFrom(self.ptr, int64(self.getLen())))
}

func (self *Node) getAny() interface{} {
	return *((*interface{})(self.ptr))
}

func (self *Node) getInt64() int64 {
	return *(*int64)(unsafe.Pointer(&self.val))
}

func (self *Node) getUint64() uint64 {
	return self.val
}

func (self *Node) getFloat64() float64 {
	return *(*float64)(unsafe.Pointer(&self.val))
}

func (self *Node) getElem(i int) *Node {
	p := (*linkedNodes)(self.ptr)
	if l := p.Len(); l != self.getLen() {
		// some nodes got unset, iterate to skip them
		for j:=0; j<l; j++ {
			v := p.At(j)
			if v.Exists() {
				i--
			}
			if i < 0 {
				return v
			}
		}
		return nil
    }
    return p.At(i)
}

func (self *Node) getPair(i int) *Pair {
	p := (*linkedPairs)(self.ptr)
	if l := p.Len(); l != self.getLen() {
		// some nodes got unset, iterate to skip them
		for j:=0; j<l; j++ {
			v := p.At(j)
			if v != nil && v.Value.Exists() {
				i--
			}
			if i < 0 {
				return v
			}
		}
		return nil
	} 
    return p.At(i)
}

// loadRaw will parse the raw value and put it in self.
func (self *Node) loadRaw() error {
	node, err := parseLazy(self.getRaw(), nil)
	if err != nil {
		return err
	}
	*self = node
	return nil
}


/***************** Factory Helper ***********************/

var (
    nullNode  = Node{typ: _V_NULL}
    trueNode  = Node{typ: _V_TRUE}
    falseNode = Node{typ: _V_FALSE}
)

func newArray(v *linkedNodes) Node {
	println("len is , ", v.Len())
    return Node{
        typ: _V_ARRAY,
        val: uint64(v.Len()),
        ptr: unsafe.Pointer(v),
    }
}

func newObject(v *linkedPairs) Node {
    return Node{
		typ: _V_OBJECT,
        val: uint64(v.Len()),
        ptr: unsafe.Pointer(v),
    }
}


// Note: not validate the input json, only used internal
func newRawNodeUnsafe(json string, hasEsc bool) Node {
	var typ valueType
	if !hasEsc {
		typ = _V_RAW
	} else {
		typ = _V_RAW_ESC
	}
	typ |= typeJumpTable[json[0]]
	
	return Node {
		typ: typ,
		val: uint64(len(json)),
		ptr: unsafe.Pointer(rt.StrPtr(json)),
	}
}

var typeJumpTable = [256]valueType{
    '"' : _V_STRING,
    '-' : _V_NUMBER,
    '0' : _V_NUMBER,
    '1' : _V_NUMBER,
    '2' : _V_NUMBER,
    '3' : _V_NUMBER,
    '4' : _V_NUMBER,
    '5' : _V_NUMBER,
    '6' : _V_NUMBER,
    '7' : _V_NUMBER,
    '8' : _V_NUMBER,
    '9' : _V_NUMBER,
    '[' : _V_ARRAY,
    'f' : _V_FALSE,
    'n' : _V_NULL,
    't' : _V_TRUE,
    '{' : _V_OBJECT,
}