package ast

import (
	"encoding/json"
	"strings"
	"unsafe"

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
	_F_MUT = types.Flag(1<<2) // mutated
	_F_RAW = types.Flag(1<<1) // raw json
	_F_KEY = types.Flag(1<<1) // string
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

func (self *Node) should2(t1 types.Type, t2 types.Type) error {
    if err := self.Error(); err != "" {
        return self
    }
    if  self.Kind != t2 && self.Kind != t2 {
        return ErrUnsupportType
    }
    return nil
}

func (n *Node) get(key string) Node  {
	if err := n.should(types.T_OBJECT); err != nil {
		return newError(err)
	}
	t, err := n.objAt(key)
	if err != nil {
		return newError(err)
	}
	return n.sub(*t)
}

func (n *Node) index(key int) Node  {
	if err := n.should(types.T_ARRAY); err != nil {
		return newError(err)
	}
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

func (self *Node) SetByPath(val Node, path ...interface{}) (bool, error) {
	if l := len(path); l == 0 {
		*self = val
		return true, nil
	} else if l == 1 {
		switch p := path[0].(type) {
		case int:
			e := self.arrSet(p, val.Kind, _F_RAW, unsafe.Pointer(&val.JSON))
			return e == nil, e 
		case string:
			e := self.objSet(p, val.Kind, _F_RAW, unsafe.Pointer(&val.JSON))
			if e == ErrNotExist {
				ee := self.objAdd(p, val.Kind, _F_RAW, unsafe.Pointer(&val.JSON))
				return false, ee
			} 
			return e == nil, e
		default:
			panic("path must be either int or string")
		}
	} else {
		parser := NewParser(self.JSON)
		start, err := parser.skip(path...)
		if err == ErrNotExist {
			// todo: make JSON and insert
		} else if err != nil {
			return false, err
		}
		sb := make([]byte, 0, start+len(val.JSON)+(len(self.JSON)-parser.pos))
		sb = append(sb, self.JSON[:start]...)
		sb = append(sb, val.JSON...)
		sb = append(sb, self.JSON[parser.pos:]...)
		return true, nil
	}
}

/***************** Cast APIs ***********************/

func (n *Node) Int64() (int64, error) {
	if err := n.should(types.T_NUMBER); err != nil {
		return 0, err
	}
	return n.toInt64()
}

func (n *Node) toInt64() (int64, error) {
	return json.Number(n.JSON).Int64()
}


/***************** Set APIs ***********************/