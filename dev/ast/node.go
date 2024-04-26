package ast

import (
	"encoding/json"

	"github.com/bytedance/sonic/encoder"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)

// 0..4 bits for basic type
const (
	_V_NONE  = 0
	_V_ERROR  = 0
)

type valueType uint8



type Pair struct {
    Key   string
    Value Node
}

/***************** Check APIs ***********************/

// Exists returns false only if the self is nil or empty node V_NONE
func (self *Node) Exists() bool {
	return self.Valid() && self.Kind != _V_NONE
}

// Valid reports if self is NOT V_ERROR or nil
func (self *Node) Valid() bool {
	if self == nil {
		return false
	}
	return self.Kind != _V_ERROR
}

// Check checks if the node itself is valid, and return:
//   - ErrNotExist If the node is nil
//   - Its underlying error If the node is V_ERROR
func (self *Node) Check() error {
	if self == nil {
		return ErrNotExist
	} else if self.Kind == _V_ERROR {
		return self
	} else {
		return nil
	}
}

// func (self *Node) IsRaw() bool{
// 	return self != nil && (self.Flag & _F_RAW != 0)
// }

const (
	_F_MUT = types.Flag(1<<2)
	// _F_RAW = types.Flag(1<<1)
)


func (self *Node) IsMut() bool{
	return self != nil && len(self.mut) != 0
}

/***************** New APIs ***********************/

var (
    nullNode  = types.NewNode("null", false)
    trueNode  = types.NewNode("true", false)
    falseNode = types.NewNode("false", false)
)

// NewRaw creates a node of raw json.
// If the input json is invalid, NewRaw returns a error Node.
func NewRaw(json string) Node {
	n, e := NewParser(json).Parse()
	if e != nil {
		return newError(e)
	}
	return n
}

// NewAny creates a node of type V_ANY if any's type isn't Node or *Node
func NewAny(val interface{}) Node {
    js, err := encoder.Encode(val, 0)
	if err != nil {
		return newError(err)
	}
	return NewRaw(rt.Mem2Str(js))
}

// NewNull creates a node of type V_NULL
func NewNull() Node {
    return Node{nullNode, nil}
}

// NewBool creates a node of type bool:
//  If v is true, returns V_TRUE node
//  If v is false, returns V_FALSE node
func NewBool(v bool) Node {
   if v {
		return Node{trueNode, nil}
   } else {
		return Node{falseNode, nil}
   }
}

// NewNumber creates a json.Number node
// v must be a decimal string complying with RFC8259
func NewNumber(v json.Number) Node {
    return newRawNodeUnsafe(string(v), false)
}

// NewString creates a node of type V_STRING. 
// v is considered to be a valid UTF-8 string, which means it won't be validated and unescaped.
// when the node is encoded to json, v will be escaped.
func NewString(v string) Node {
	s := encoder.Quote(v)
	esc := len(s) > len(v) + 2
	return newRawNodeUnsafe(s, esc)
}

func NewInt(v int) Node {
	s, err := encoder.Encode(v, 0)
	if err != nil {
		return newError(err)
	}
	return newRawNodeUnsafe(rt.Mem2Str(s), false)
}

func NewFloat(v float64) Node {
	s, err := encoder.Encode(v, 0)
	if err != nil {
		return newError(err)
	}
	return newRawNodeUnsafe(rt.Mem2Str(s), false)
}

// NewBytes encodes given src with Base64 (RFC 4648), and creates a node of type V_STRING.
func NewBytes(src []byte) Node {
    if len(src) == 0 {
        panic("empty src bytes")
    }
    out := encodeBase64(src)
    return NewString(out)
}

func (self *Node) should(t types.Type) error {
    if err := self.Error(); err != "" {
        return self
    }
    if  self.Kind != t {
        return ErrUnsupportType
    }
    return nil
}

func (n *Node) get(key string) (Node)  {
	t, err := n.objAt(key)
	if err != nil {
		return newError(err)
	}
	return n.sub(*t)
}

func (n *Node) index(key int) (Node)  {
	t := n.arrAt(key)
	if t == nil {
		return emptyNode
	}
	return n.sub(*t)
}


func (self *Node) GetByPath(path ...interface{}) Node {
	if l := len(path); l == 0 {
		return *self
	} else if l == 1 {
		switch p := path[0].(type) {
		case int:
			return self.index(p)
		case string:
			return self.get(p)
		default:
			panic("path must be either int or string")
		}
	} else {
		n, err := NewParser(self.JSON).getByPath(path...)
		if err != nil {
			return newError(err)
		}
		return n
	}
}

/***************** Cast APIs ***********************/




/***************** Set APIs ***********************/