package ast

import (
	"sync"
	"unsafe"
	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/native/types"
)

type token uint64

// a internal node parsed from native C.
type node struct {
	kind token
	tape []token
	json string
}

var tokenPool = sync.Pool{
    New: func () interface{} {
        return make([]token, 0, 1024)
    },
}

// encoding 64-bit token as follows:
const (
	/* 63 ~ 64 bits */
	_KLITERAL 	token	= (0<<62)
	_KRAW 		token	= (1<<62)
	_KRAW_ESC 	token 	= (2<<62)
	_KTYPED   	token 	= (3<<62)
	_MASK_63_64 token  	= (3<<62)

	/* 33 ~ 62 bits for raw/obj/arr length */
	_MAX_LEN 	token 	= ((1 << 30) - 1)

	/* 1 ~ 32 bits for offset or details */
	_MAX_OFF 	token 	= ((1 << 32) - 1)

	_KNULL 		token 	= _KLITERAL | 0
	_KTRUE 		token 	= _KLITERAL | 1
	_KFALSE 	token 	= _KLITERAL | 2

	_KOBJ_TYPE 	token	= _KTYPED 	| 0
	_KARR_TYPE 	token 	= _KTYPED 	| 1

	/* specific error code */
	_MUST_RETRY = 0x12345
)

func (t token) isRaw() bool {
	return (t & _MASK_63_64) == _KRAW || (t & _MASK_63_64) == _KRAW_ESC
}

func (t token) hasEsc() bool {
	return (t & _MASK_63_64) == _KRAW_ESC
}

func (t token) offset() int64 {
	return int64(t & _MAX_OFF)
}

func (t token) len() int64 {
	return int64((t  >> 32) & _MAX_LEN)
}

func (n *node) objAt(i int) (Pair, error)  {
	key, err := n.tape[i * 2].str(n.json)
	if err != nil {
		return Pair{}, err
	}

	node := n.tape[i * 2 + 1]
	val := node.raw(n.json)
	return Pair{key, newRawNodeUnsafe(val, node.hasEsc())}, nil
}

func (n *node) arrAt(i int) Node {
	node := n.tape[i * 2 + 1]
	val := node.raw(n.json)
	return newRawNodeUnsafe(val, node.hasEsc())
}

func (t token) peek(json string) byte {
	return json[t.offset()]
}

func (t token) raw(json string) string {
	return json[t.offset():t.offset() + t.len()]
}

func (t token) str(json string) (string, error) {
	raw := t.raw(json)
	if !t.hasEsc(){
		// remove the quotes
		return raw[1: len(raw) - 1], nil
	}
	
	s, err := unquote(raw)
	if err != 0 {
		return "", makeSyntaxError(json, int(t.offset()), err.Message())
	} else {
		return s, nil
	}
}

func makeSyntaxError(json string, p int, msg string) SyntaxError {
	return SyntaxError{
		Pos : p,
		Src : json,
		Msg: msg,
	}
}

// TODO: use flags to make, if is primitives
func parseLazy(json string, path *[]interface{}) (Node, error) {
	node := &node {
		kind: _KNULL,
		tape: tokenPool.Get().([]token),
		json: json,
	}

	defer tokenPool.Put(node.tape)

	/* parse into inner node */
	r, p := 0, 0
	for {
		p = 0
		r = native.ParseLazy(&json, &p, ((*native.Node)(unsafe.Pointer(node))), path)
		if r == - _MUST_RETRY {
			node.tape = make([]token, 0, 2 * cap(node.tape))
		} else {
			break
		}
	}

	/* check errors */
	if r < 0 {
		return  Node{},  makeSyntaxError(json, p, types.ParsingError(-r).Message())
	}

	println("json xxx is ", json)
	/* convert to ast.Node */
	var err error
	switch node.kind {
        case _KNULL    : return nullNode, nil
        case _KTRUE    : return trueNode, nil
        case _KFALSE   : return falseNode, nil
		case _KOBJ_TYPE : {
			pairs := make([]Pair, len(node.tape) / 2)
			for i := 0; i < len(pairs); i++ {
				pairs[i], err = node.objAt(i)
				if err != nil {
					return Node{}, err
				}
			}
			s := new(linkedPairs)
			s.FromSlice(pairs)
			return newObject(s), nil
		}
		case _KARR_TYPE: {
			println("parse arrt lay , json is ", json)
			nodes := make([]Node, len(node.tape))
			for i := 0; i < len(nodes); i++ {
				nodes[i] = node.arrAt(i)
			}
			s := new(linkedNodes)
			s.FromSlice(nodes)
			return newArray(s), nil
		}
		default:  {
			return newRawNodeUnsafe(node.json, node.kind.hasEsc()), nil
		}
	}
}

func unquoteFast(json string) string {
	return json[1:len(json) - 1]
}
