package ast

import (
	"unsafe"

	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/native/types"
)

// a internal node parsed from native C.
type Node struct {
	node types.Node
	mut []unsafe.Pointer // (nil)
}

func (n *Node) arrSet(i int, typ types.Type, flag types.Flag,  val unsafe.Pointer) error {
	t := n.arrAt(i)
	if t == nil {
		return ErrNotExist
	}
	l := len(n.mut)
	*t = types.Token{
		Kind: typ,
		Flag: _F_MUT | flag,
		Off: uint32(l),
	}
	n.mut = append(n.mut, val)
	n.node.Flag |= _F_MUT
	return nil
}

func (n *Node) arrAdd(typ types.Type, flag types.Flag, val unsafe.Pointer) error {
	l := len(n.mut)
	v := types.Token{
		Kind: typ,
		Flag: _F_MUT | flag,
		Off: uint32(l),
	}

	n.mut = append(n.mut, val)
	n.node.Flag |= _F_MUT
	
	n.node.Kids = append(n.node.Kids, v)
	return nil
}

func (n *Node) arrDel(i int) error {
	t := n.arrAt(i)
	if t == nil {
		return ErrNotExist
	}
	var right []types.Token
	if i < len(n.node.Kids) - 1 {
		right = n.node.Kids[i+1:]
	}
	n.node.Kids = append(n.node.Kids[:i], right...)
	if t.Flag & _F_MUT != 0 {
		x := int(t.Off)
		var right []unsafe.Pointer
		if x < len(n.mut) - 1 {
			right = n.mut[x+1:]
		}
		n.mut = append(n.mut[:x], right...)
	}
	return nil
}

func (n *Node) objSet(key string, typ types.Type, flag types.Flag, val unsafe.Pointer) error {
	_, t, err := n.objAt(key)
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
	n.node.Flag |= _F_MUT
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
	n.node.Flag |= _F_MUT
	
	n.node.Kids = append(n.node.Kids, k)
	n.node.Kids = append(n.node.Kids, v)
	return nil
}

func (n *Node) objAt(key string) (int, *types.Token, error)  {
	for i := 0; i<len(n.node.Kids); i+=2 {
		k, err := n.str(n.node.Kids[i])
		if err != nil {
			return -1, nil, err
		}
		if k == key {
			return i, &n.node.Kids[i+1], nil
		}
	}
	return -1, nil, ErrNotExist
}

func (n *Node) objDel(key string) error {
	i, t, err := n.objAt(key)
	if err != nil {
		return err
	}
	if t == nil {
		return ErrNotExist
	}
	var right []types.Token
	if i < len(n.node.Kids) - 2 {
		right = n.node.Kids[i+2:]
	}
	n.node.Kids = append(n.node.Kids[:i], right...)
	if t.Flag & _F_MUT != 0 {
		x := int(t.Off)
		var right []unsafe.Pointer
		if x < len(n.mut) - 2 {
			right = n.mut[x+2:]
		}
		n.mut = append(n.mut[:x], right...)
	}
	return nil
}

// This will convert a token to Node
//   - scalar type, directly slice original string
//   - array/object, parse to Node for one layer
//   - mut type, use unsafe.Pointer casting
// TODO: handle mut token
func (n *Node) sub(t types.Token) Node {
	if t.Flag & _F_MUT == 0 {
		return newRawNodeUnsafe(t.Raw(n.node.JSON), t.Flag.IsEsc())
	} else {
		panic("not implement!")
	}
}

func (n *Node) arrAt(i int) *types.Token {
	if i >= len(n.node.Kids) {
		return nil
	}
	return &n.node.Kids[i]
}

func (n *Node) json(t types.Token) string {
	return t.Raw(n.node.JSON)
}

func (n *Node) str(t types.Token) (string, error) {
	raw := n.json(t)
	if !t.Flag.IsEsc(){
		// remove the quotes
		return raw[1: len(raw) - 1], nil
	}
	
	s, err := unquote(raw)
	if err != 0 {
		return "", makeSyntaxError(raw, int(t.Off), err.Message())
	} else {
		return s, nil
	}
}


// TODO: use flags to make, if is primitives
func parseLazy(json string, path *[]interface{}) (Node, error) {
	// TODO: got real PC of biz caller
	node := Node{}
	node.node.Kids = make([]types.Token, 0, types.PredictTokenSize())

	/* parse into inner node */
	r, p := 0, 0
	for {
		p = 0
		r = native.ParseLazy(&json, &p, &node.node, path)
		if r == -types.MUST_RETRY {
			node.node.Grow()
		} else {
			break
		}
	}

	/* check errors */
	if r < 0 {
		return Node{},  makeSyntaxError(json, p, types.ParsingError(-r).Message())
	}

	node.node.JSON = json[r:p]
	types.RecordTokenSize(int64(len(node.node.Kids)))
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
