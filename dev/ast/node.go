package ast

import (
	"encoding/json"
	"unsafe"

	"github.com/bytedance/sonic/encoder"
	"github.com/bytedance/sonic/internal/encoder/alg"
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
	return self.Valid() && self.node.Kind != _V_NONE
}

// Valid reports if self is NOT V_ERROR or nil
func (self *Node) Valid() bool {
	if self == nil {
		return false
	}
	return self.node.Kind != _V_ERROR
}

// Check checks if the node itself is valid, and return:
//   - ErrNotExist If the node is nil
//   - Its underlying error If the node is V_ERROR
func (self *Node) Check() error {
	if self == nil {
		return ErrNotExist
	} else if self.node.Kind == _V_ERROR {
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
    if  self.node.Kind != t {
        return ErrUnsupportType
    }
    return nil
}

func (self *Node) should2(t1 types.Type, t2 types.Type) error {
    if err := self.Error(); err != "" {
        return self
    }
    if  self.node.Kind != t1 && self.node.Kind != t2 {
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
		n, err := NewParser(self.node.JSON).GetByPath(path...)
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
		// for one layer set
		switch p := path[0].(type) {
		case int:
			e := self.arrSet(p, val.node.Kind, _F_RAW, unsafe.Pointer(&val.node.JSON))
			return e == nil, e 
		case string:
			e := self.objSet(p, val.node.Kind, _F_RAW, unsafe.Pointer(&val.node.JSON))
			if e == ErrNotExist {
				ee := self.objAdd(p, val.node.Kind, _F_RAW, unsafe.Pointer(&val.node.JSON))
				return false, ee
			} 
			return e == nil, e
		default:
			panic("path must be either int or string")
		}
	} else {
		// multi layers set
		p := NewParser(self.node.JSON)
		var err types.ParsingError
		var idx int
		var start int
		for i, k := range path {
			if id, ok := k.(int); ok {
				if start, err = p.locate(id); err != 0 {
					return false, makeSyntaxError(self.node.JSON, p.pos, err.Message())
				}
			} else if key, ok := k.(string); ok {
				if start, err = p.locate(key); err != 0 {
					// for object, we allow insert non-exist key
					if err != types.ERR_NOT_FOUND {
						return false, makeSyntaxError(self.node.JSON, p.pos, err.Message())
					} else {
						idx = i
						break
					}
				}
			} else {
				panic("path must be either int or string")
			}
		}
		var b []byte
		if err == types.ERR_NOT_FOUND {
			s := p.pos - 1
			for ; s >= 0 && isSpace(self.node.JSON[s]); s-- {
			}
			empty := (self.node.JSON[s] == '[' || self.node.JSON[s] == '{')
			size := len(self.node.JSON) + len(val.node.JSON) + 8*(len(path))
			b = make([]byte, 0, size)
			s = s + 1
			b = append(b, self.node.JSON[:s]...)
			if !empty {
				b = append(b, ","...)
			}
			// creat new nodes on path
			var err error
			b, err = makePathAndValue(b, path[idx:], false, val)
			if err != nil {
				return true, err
			}
		} else if err != 0 {
			return false, makeSyntaxError(self.node.JSON, p.pos, err.Message())
		} else {
			b = make([]byte, 0, start+len(val.node.JSON)+(len(self.node.JSON)-p.pos))
			b = append(b, self.node.JSON[:start]...)
			b = append(b, val.node.JSON...)
			b = append(b, self.node.JSON[p.pos:]...)
		}
		// refrest the node
		p.src = rt.Mem2Str(b)
		p.pos = 0
		node, ee := p.Parse()
		if ee != nil {
			return true, ee
		}
		*self = node
		return true, nil
	}
}

// [2,"a"],1 => {"a":1}
// ["a",2],1  => "a":[1]
func makePathAndValue(b []byte, path []interface{}, allowAppend bool, val Node) ([]byte, error) {
	for i, k := range path {
		if key, ok := k.(string); ok {
			b = alg.Quote(b, key, false)
			b = append(b, ":"...)
		}
		if i == len(path)-1 {
			b = append(b, val.node.JSON...)
			break
		}
		n := path[i+1]
		if _, ok := n.(int); ok {
			if !allowAppend {
				return nil, ErrOutOfRange
			}
			b = append(b, "["...)
		} else if _, ok := n.(string); ok {
			b = append(b, `{`...)
		} else {
			panic("path must be either int or string")
		}
	}
	for i := len(path) - 1; i >= 1; i-- {
		k := path[i]
		if _, ok := k.(int); ok {
			b = append(b, "]"...)
		} else if _, ok := k.(string); ok {
			b = append(b, `}`...)
		}
	}
	return b, nil
}

/***************** Cast APIs ***********************/

func (n *Node) Int64() (int64, error) {
	if err := n.should(types.T_NUMBER); err != nil {
		return 0, err
	}
	return n.toInt64()
}

func (n *Node) toInt64() (int64, error) {
	return json.Number(n.node.JSON).Int64()
}


/***************** Set APIs ***********************/