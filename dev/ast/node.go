package ast

import (
	"encoding/json"
	"unsafe"

	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/encoder"
	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)

/***************** Check APIs ***********************/

// Exists returns false only if the self is nil or empty node V_NONE
func (self *Node) Exists() bool {
	return self.Valid() && self.node.Kind != V_NONE
}

// Valid reports if self is NOT V_ERROR or nil
func (self *Node) Valid() bool {
	if self == nil {
		return false
	}
	return self.node.Kind != V_ERROR
}

// Check checks if the node itself is valid, and return:
//   - ErrNotExist If the node is nil
//   - Its underlying error If the node is V_ERROR
func (self *Node) Check() error {
	if self == nil {
		return ErrNotExist
	} else if self.node.Kind == V_ERROR {
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

func NewInt64(v int64) Node {
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

const (
    V_NONE   = 0
    V_ERROR  = 1
    V_NULL   = int(types.T_NULL)
    V_TRUE   = int(types.T_TRUE)
    V_FALSE  = int(types.T_FALSE)
    V_ARRAY  = int(types.T_ARRAY)
    V_OBJECT = int(types.T_OBJECT)
    V_STRING = int(types.T_STRING)
    V_NUMBER = int(types.T_NUMBER)
)

// Type returns json type represented by the node
// It will be one of belows:
//
//	V_NONE   = 0 (empty node)
//	V_ERROR  = 1 (something wrong)
//	V_NULL   = 2 (json value `null`)
//	V_TRUE   = 3 (json value `true`)
//	V_FALSE  = 4 (json value `false`)
//	V_ARRAY  = 5 (json value array)
//	V_OBJECT = 6 (json value object)
//	V_STRING = 7 (json value string)
//	V_NUMBER = 8 (json value number )
func (self Node) Type() int {
	return int(self.node.Kind)
}

func (n *Node) Raw() (string, error) {
	if n.Error() != "" {
		return "", n
	}
	return n.node.JSON, nil
}

func (n *Node) Bool() (bool, error) {
	switch n.node.Kind {
		case types.T_FALSE: return false, nil
		case types.T_TRUE: return true, nil
		default: return false, ErrUnsupportType
	}
}

func (n *Node) Int64() (int64, error) {
	if err := n.should(types.T_NUMBER); err != nil {
		return 0, err
	}
	var st types.JsonState
	e := native.Value(unsafe.Pointer(&n.node.JSON), len(n.node.JSON), 0, &st, 0)
	if e < 0 {
		return 0, makeSyntaxError(n.node.JSON, 0, types.ParsingError(e).Message())
	}
	switch st.Vt {
		case types.V_INTEGER: return st.Iv, nil
		default: return 0, ErrUnsupportType
	}
}

func (n *Node) Float64() (float64, error) {
	if err := n.should(types.T_NUMBER); err != nil {
		return 0, err
	}
	var st types.JsonState
	e := native.Value(unsafe.Pointer(&n.node.JSON), len(n.node.JSON), 0, &st, 0)
	if e < 0 {
		return 0, makeSyntaxError(n.node.JSON, 0, types.ParsingError(e).Message())
	}
	switch st.Vt {
		case types.V_INTEGER: return float64(st.Iv), nil
		case types.V_DOUBLE: return st.Dv, nil
		default: return 0, ErrUnsupportType
	}
}

func (n *Node) Number() (json.Number, error) {
	if err := n.should(types.T_NUMBER); err != nil {
		return "", err
	}
	return json.Number(n.node.JSON), nil
}

func (n *Node) String() (string, error) {
	if err := n.should(types.T_STRING); err != nil {
		return "", err
	}
	if !n.node.Flag.IsEsc() {
		return n.node.JSON[1: len(n.node.JSON) - 1], nil
	}
	s, err := unquote(n.node.JSON)
	if err != 0 {
		return "", makeSyntaxError(n.node.JSON, 0, err.Message())
	} else {
		return s, nil
	}
}

func (n *Node) Array(buf *[]Node) error {
	if err := n.should(types.T_ARRAY); err != nil {
		return err
	}
	l := len(n.node.Kids)
	ol := len(*buf)
	if cap(*buf) - ol < l {
		tmp := make([]Node, ol+l)
		copy(tmp, *buf)
		*buf = tmp[:ol]
	}
	for _, t := range n.node.Kids {
		*buf = append(*buf, n.sub(t))
	}
	return nil
}

func (n *Node) Map(buf map[string]Node) error {
	if err := n.should(types.T_OBJECT); err != nil {
		return err
	}
	for i := 0; i<len(n.node.Kids);  {
		t := n.node.Kids[i];
		key, err := n.str(t)
		if err != nil {
			return err
		}
		tt := n.node.Kids[i+1]
		val := n.sub(tt)
		buf[key] = val
		i += 2
	}
	return nil
}

func (n *Node) InterfaceUseNode() (interface{}, error) {
	switch n.node.Kind {
	case types.T_NULL:
		return nil, nil
	case types.T_FALSE:
		return false, nil
	case types.T_TRUE:
		return true, nil
	case types.T_STRING:
		return n.String()
	case types.T_NUMBER:
		return n.Number()
	case types.T_ARRAY:
		buf := make([]Node, 0, len(n.node.Kids))
		err := n.Array(&buf)
		return buf, err
	case types.T_OBJECT:
		buf := make(map[string]Node, len(n.node.Kids)/2)
		err := n.Map(buf)
		return buf, err
	case V_ERROR:
		return nil, n
	case V_NONE:
		return nil, ErrNotExist
	default:
		return nil, ErrUnsupportType
	}
}

func (n *Node) InterfaceUseGoPrimitive(opts decoder.Options) (interface{}, error) {
	switch n.node.Kind {
	case types.T_NULL:
		return nil, nil
	case types.T_FALSE:
		return false, nil
	case types.T_TRUE:
		return true, nil
	case types.T_STRING:
		return n.String()
	case types.T_NUMBER:
		if opts & decoder.OptionUseNumber != 0 {
			return json.Number(n.node.JSON), nil
		} else if opts & decoder.OptionUseInt64 != 0 {
			if iv, err := n.Int64(); err == nil {
				return iv, nil
			}
		}
		return n.Float64()
	case types.T_ARRAY:
		buf := make([]interface{}, 0, len(n.node.Kids))
		for i, tok := range n.node.Kids {
			js := tok.Raw(n.node.JSON)
			dc := decoder.NewDecoder(js)
			dc.SetOptions(opts)
			if err := dc.Decode(&buf[i]); err != nil {
				return nil, err
			}
		}
		return buf, nil
	case types.T_OBJECT:
		buf := make(map[string]interface{}, len(n.node.Kids)/2)
		for i:=0; i<len(n.node.Kids); i += 2 {
			key, err := n.str(n.node.Kids[i])
			if err != nil {
				return nil, err
			}
			var val interface{}
			js := n.node.Kids[i+1].Raw(n.node.JSON)
			dc := decoder.NewDecoder(js)
			dc.SetOptions(opts)
			if err := dc.Decode(&val); err != nil {
				return nil, err
			}
			buf[key] = val
		}
		return buf, nil
	case V_ERROR:
		return nil, n
	case V_NONE:
		return nil, ErrNotExist
	default:
		return nil, ErrUnsupportType
	}
}

func (n *Node) ForEachKV(scanner func(key string, elem Node) bool) error {
	if err := n.should(types.T_OBJECT); err != nil {
		return err
	}
	for i:=0; i<len(n.node.Kids); i+=2 {
		key, err := n.str(n.node.Kids[i])
		if err != nil {
			return err
		}
		val := n.sub(n.node.Kids[i+1])
		if !scanner(key, val) {
			return nil
		}
	}
	return nil
}

func (n *Node) ForEachElem(scanner func(index int, elem Node) bool) error {
	if err := n.should(types.T_OBJECT); err != nil {
		return err
	}
	for i, t := range n.node.Kids {
		elem := n.sub(t)
		if !scanner(i, elem) {
			return nil
		}
	}
	return nil
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

var EstimatedInsertedPathCharSize = 8

func (self *Node) SetByPath(allowArrayAppend bool, val Node, path ...interface{}) (bool, error) {
	if l := len(path); l == 0 {
		*self = val
		return true, nil
	} else if l == 1 {
		// for one layer set
		switch p := path[0].(type) {
		case int:
			e := self.arrSet(p, val.node.Kind, _F_RAW, unsafe.Pointer(&val.node.JSON))
			if e == ErrNotExist && allowArrayAppend {
				ee := self.arrAdd(val.node.Kind, _F_RAW, unsafe.Pointer(&val.node.JSON))
				return false, ee
			}
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
		var missing int
		var start int
		for i, k := range path {
			if id, ok := k.(int); ok {
				if start, err = p.locate(id); err != 0 {
					if err != types.ERR_NOT_FOUND {
						return false, makeSyntaxError(self.node.JSON, p.pos, err.Message())
					} else {
						missing = i
						break
					}
				}
			} else if key, ok := k.(string); ok {
				if start, err = p.locate(key); err != 0 {
					// for object, we allow insert non-exist key
					if err != types.ERR_NOT_FOUND {
						return false, makeSyntaxError(self.node.JSON, p.pos, err.Message())
					} else {
						missing = i
						break
					}
				}
			} else {
				panic("path must be either int or string")
			}
		}
		var b []byte
		if err == types.ERR_NOT_FOUND {
			// NOTICE: pos must stop at '}' or ']'
			s, empty := backward(self.node.JSON, p.pos)
			path := path[missing:]
			size := len(self.node.JSON) + len(val.node.JSON) + EstimatedInsertedPathCharSize*(len(path))
			b = make([]byte, 0, size)
			b = append(b, self.node.JSON[:s]...)
			if !empty {
				b = append(b, ","...)
			}
			// creat new nodes on path
			var err error
			b, err = makePathAndValue(b, path, allowArrayAppend, val)
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

func 