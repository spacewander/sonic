package ast

import (
	"unsafe"

	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/native/types"
)

// a internal node parsed from native C.
type Node struct {
	types.Node
	mut []unsafe.Pointer // (nil)
}

func (n *Node) arrSet(i int, typ types.Type, flag types.Flag,  val unsafe.Pointer) error {
	t := n.arrAt(i)
	if t == nil {
		return ErrOutOfRange
	}
	l := len(n.mut)
	*t = types.Token{
		Kind: typ,
		Flag: _F_MUT | flag,
		Off: uint32(l),
	}
	n.mut = append(n.mut, val)
	n.Flag |= _F_MUT
	return nil
}

func (n *Node) objSet(key string, typ types.Type, flag types.Flag, val unsafe.Pointer) error {
	t, err := n.objAt(key)
	if err != nil {
		return err
	}
	l := len(n.mut)
	*t = types.Token{
		Kind: typ,
		Flag: _F_MUT | flag,
		Off: uint32(l),
	}
	n.mut = append(n.mut, val)
	n.Flag |= _F_MUT
	return nil
}

func (n *Node) objAdd(key string, typ types.Type, flag types.Flag, val unsafe.Pointer) error {
	l := len(n.mut)
	k := types.Token{
		Kind: types.T_STRING,
		Flag: _F_KEY,
		Off: uint32(l),
	}
	v := types.Token{
		Kind: typ,
		Flag: _F_MUT | flag,
		Off: uint32(l+1),
	}

	n.mut = append(n.mut, unsafe.Pointer(&key))
	n.mut = append(n.mut, val)
	n.Flag |= _F_MUT
	
	n.Node.Kids = append(n.Node.Kids, k)
	n.Node.Kids = append(n.Node.Kids, v)
	return nil
}

func (n *Node) objAt(key string) (*types.Token, error)  {
	for i := 0; i<len(n.Kids)/2; i++ {
		k, err := strForToken(n.Kids[i * 2], n.JSON)
		if err != nil {
			return nil, err
		}
		if k == key {
			return &n.Kids[i * 2+1], nil
		}
	}
	return nil, ErrNotExist
}

// This will convert a token to Node
//   - scalar type, directly slice original string
//   - array/object, parse to Node for one layer
//   - mut type, use unsafe.Pointer casting
// TODO: handle mut token
func (n *Node) sub(t types.Token) Node {
	if t.Flag & _F_MUT == 0 {
		return newRawNodeUnsafe(t.Raw(n.JSON), t.Flag.IsEsc())
	} else {
		panic("not implement!")
	}
}

func (n *Node) arrAt(i int) *types.Token {
	if i >= len(n.Kids) {
		return nil
	}
	return &n.Kids[i]
}

func strForToken(t types.Token, json string) (string, error) {
	raw := t.Raw(json)
	if !t.Flag.IsEsc(){
		// remove the quotes
		return raw[1: len(raw) - 1], nil
	}
	
	s, err := unquote(raw)
	if err != 0 {
		return "", makeSyntaxError(json, int(t.Off), err.Message())
	} else {
		return s, nil
	}
}

func makeSyntaxError(json string, p int, msg string) decoder.SyntaxError {
	return decoder.SyntaxError{
		Pos : p,
		Src : json,
		Msg: msg,
	}
}

// TODO: use flags to make, if is primitives
func parseLazy(json string, path *[]interface{}) (Node, error) {
	// TODO: got real PC of biz caller
	node := Node{}
	node.Node.Kids = make([]types.Token, 0, types.PredictTokenSize())

	/* parse into inner node */
	r, p := 0, 0
	for {
		p = 0
		r = native.ParseLazy(&json, &p, &node.Node, path)
		if r == -types.MUST_RETRY {
			node.Node.Grow()
		} else {
			break
		}
	}

	/* check errors */
	if r < 0 {
		return Node{},  makeSyntaxError(json, p, types.ParsingError(-r).Message())
	}

	node.JSON = json
	types.RecordTokenSize(int64(len(node.Node.Kids)))
	return node, nil

	// // println("json xxx is ", json)
	// /* convert to ast.Node */
	// var err error
	// switch node.kind.kind {
    //     case _KNULL    : return nullNode, nil
    //     case _KTRUE    : return trueNode, nil
    //     case _KFALSE   : return falseNode, nil
	// 	case _KOBJ_TYPE : {
	// 		pairs := make([]Pair, len(node.tape) / 2)
	// 		for i := 0; i < len(pairs); i++ {
	// 			pairs[i], err = node.objAt(i)
	// 			if err != nil {
	// 				return Node{}, err
	// 			}
	// 		}
	// 		s := new(linkedPairs)
	// 		s.FromSlice(pairs)
	// 		return newObject(s), nil
	// 	}
	// 	case _KARR_TYPE: {
	// 		println("parse arrt lay , json is ", json)
	// 		nodes := make([]Node, len(node.tape))
	// 		for i := 0; i < len(nodes); i++ {
	// 			nodes[i] = node.arrAt(i)
	// 		}
	// 		s := new(linkedNodes)
	// 		s.FromSlice(nodes)
	// 		return newArray(s), nil
	// 	}
	// 	default:  {
	// 		return newRawNodeUnsafe(node.json, node.kind.hasEsc()), nil
	// 	}
	// }
}


// Note: not validate the input json, only used internal
func newRawNodeUnsafe(json string, hasEsc bool) Node {
	n := types.NewNode(json, hasEsc)
	if !n.Kind.IsComplex() {
		return Node{n, nil}
	}
	return NewRaw(json)
}
