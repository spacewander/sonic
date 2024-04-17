package ast

import (
	"unsafe"

	"github.com/bytedance/sonic/internal/rt"
)

// 0..4 bits for basic type
const (
	_TYPE_BITS = 4
	_TYPE_MASK = 0xF

	_V_NONE  = 0
	_V_NULL  = 1
	_V_ERROR = 2
	_V_TRUE  = 3
	_V_FALSE = 4

	_V_INT64   = 5
	_V_UINT64  = 6
	_V_FLOAT64 = 7
	_V_NUMBER  = 8 // json.Number

	_V_STRING = 9
	_V_ARRAY  = 10
	_V_OBJECT = 11
	_V_ANY    = 12
)

// 4..8 bits for subtype of Raw
const (
	_RAW_BITS = 4
	_RAW_MASK = 0xF0

	_V_RAW     = 0x10
	_V_RAW_ESC = 0x20
)

type valueType uint8

type Node struct {
	typ valueType
	// val is value for scalar type or length for container type
	val uint64
	// pointer of non-scalar node
	ptr unsafe.Pointer
}

type Pair struct {
    Key   string
    Value Node
}

/***************** Check APIs ***********************/

// Exists returns false only if the self is nil or empty node V_NONE
func (self *Node) Exists() bool {
	return self.Valid() && self.getType() != _V_NONE
}

// Valid reports if self is NOT V_ERROR or nil
func (self *Node) Valid() bool {
	if self == nil {
		return false
	}
	return self.getType() != _V_ERROR
}

// Check checks if the node itself is valid, and return:
//   - ErrNotExist If the node is nil
//   - Its underlying error If the node is V_ERROR
func (self *Node) Check() error {
	if self == nil {
		return ErrNotExist
	} else if self.getType() == _V_ERROR {
		return self
	} else {
		return nil
	}
}

func (self *Node) IsRaw() bool{
	return self != nil && self.isRaw()
}


/***************** New APIs ***********************/

// NewRaw creates a node of raw json.
// If the input json is invalid, NewRaw returns a error Node.
func NewRaw(json string) Node {
	p := NewParser(json)
    start, err := p.skip()
    if err != nil {
        return *newError(err)
    }

	// TODO: FIXME, should has escaped flags
    return newRawNodeUnsafe(json[start: p.pos], true)
}

// NewAny creates a node of type V_ANY if any's type isn't Node or *Node
func NewAny(val interface{}) Node {
    switch n := val.(type) {
    case Node:
        return n
    case *Node:
        return *n
    default:
        return Node{
            typ: _V_ANY,
			val: 0,
            ptr: unsafe.Pointer(&val),
        }
    }
}

// NewNull creates a node of type V_NULL
func NewNull() Node {
    return nullNode
}

// NewBool creates a node of type bool:
//  If v is true, returns V_TRUE node
//  If v is false, returns V_FALSE node
func NewBool(v bool) Node {
   if v {
		return trueNode
   } else {
		return falseNode
   }
}

// NewNumber creates a json.Number node
// v must be a decimal string complying with RFC8259
func NewNumber(v string) Node {
    return Node{
		typ: _V_NUMBER,
       	val: uint64(len(v)),
		ptr: rt.StrPtr(v),
    }
}

// NewString creates a node of type V_STRING. 
// v is considered to be a valid UTF-8 string, which means it won't be validated and unescaped.
// when the node is encoded to json, v will be escaped.
func NewString(v string) Node {
    return Node{
		typ: _V_STRING,
       	val: uint64(len(v)),
		ptr: rt.StrPtr(v),
    }
}

func NewInt(v int) Node {
	if v >= 0 {
		return NewUint(uint(v))
	} else {
		cast := *(*uint64)(unsafe.Pointer(&v))
		return Node{
			typ: _V_INT64,
			val: cast,
		}
	}
}

func NewUint(v uint) Node {
	return Node{
		typ: _V_UINT64,
		val: uint64(v),
	}
}

func NewFloat(v float64) Node {
	cast := *(*uint64)(unsafe.Pointer(&v))
	return Node{
		typ: _V_FLOAT64,
		val: cast,
	}
}

// NewObject creates a node of type V_OBJECT,
// using v as its underlying children
func NewObject(v []Pair) Node {
    s := new(linkedPairs)
    s.FromSlice(v)
    return newObject(s)
}

// NewArray creates a node of type V_ARRAY,
// using v as its underlying children
func NewArray(v []Node) Node {
    s := new(linkedNodes)
    s.FromSlice(v)
    return newArray(s)
}

// NewBytes encodes given src with Base64 (RFC 4648), and creates a node of type V_STRING.
func NewBytes(src []byte) Node {
    if len(src) == 0 {
        panic("empty src bytes")
    }
    out := encodeBase64(src)
    return NewString(out)
}

/***************** Cast APIs ***********************/


/***************** Set APIs ***********************/