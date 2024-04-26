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


// func (n *node) SetAt(i int32, val unsafe.Pointer) {
// 	x := len(n.mut)
// 	n.tape[i] = token(_K_MUT|x)
// 	n.mut = append(n.mut, val)
// }


// func (n *Node) objAt(i int) (Pair, error)  {
// 	key, err := strForToken(n.Kids[i * 2], n.JSON)
// 	if err != nil {
// 		return Pair{}, err
// 	}

// 	node := n.Kids[i * 2 + 1]
// 	val := node.Raw(n.JSON)
// 	return Pair{key, newRawNodeUnsafe(val, node.Flag.IsEsc())}, nil
// }

// func (n *Node) arrAt(i int) Node {
// 	node := n.Kids[i * 2 + 1]
// 	val := node.Raw(n.JSON)
// 	return newRawNodeUnsafe(val, node.Flag.IsEsc())
// }


// func strForToken(t types.Token, json string) (string, error) {
// 	raw := t.Raw(json)
// 	if !t.Flag.IsEsc(){
// 		// remove the quotes
// 		return raw[1: len(raw) - 1], nil
// 	}
	
// 	s, err := unquote(raw)
// 	if err != 0 {
// 		return "", makeSyntaxError(json, int(t.Off), err.Message())
// 	} else {
// 		return s, nil
// 	}
// }

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
	ret := Node{
		Node: types.NewNode(json, hasEsc),
	}
	// if (ret.Kind == types.T_ARRAY || ret.Kind == types.T_OBJECT) && isRaw {
	// 	ret.Flag |= _F_RAW
	// }
	return ret
}
