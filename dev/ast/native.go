package ast

import (
	"github.com/bytedance/sonic/internal/native"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/davecgh/go-spew/spew"
)

const (
	_F_MUT = types.Flag(1<<2) // mutated
	_F_KEY = types.Flag(1<<3) // key string in mut
	// _F_RAW = types.Flag(1<<4) // raw json
)

func (n *Node) arrAt(i int) *types.Token {
	if i >= len(n.node.Kids) {
		return nil
	}
	return &n.node.Kids[i]
}

func (n *Node) arrSet(i int, val interface{}) error {
	t := n.arrAt(i)
	if t == nil {
		return ErrNotExist
	}
	l := len(n.mut)
	*t = types.Token{
		Kind: types.Type(V_ANY),
		Off: uint32(l),
	}
	n.mut = append(n.mut, val)

	n.node.Flag |= _F_MUT
	return nil
}

func (n *Node) arrAdd( val interface{}) error {
	l := len(n.mut)
	v := types.Token{
		Kind: types.Type(V_ANY),
		Off: uint32(l),
	}
	n.mut = append(n.mut, val)
	n.node.Kids = append(n.node.Kids, v)

	n.node.Flag |= _F_MUT
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
	if t.Kind == types.Type(V_ANY) {
		x := int(t.Off)
		var right []interface{}
		if x < len(n.mut) - 1 {
			right = n.mut[x+1:]
		}
		n.mut = append(n.mut[:x], right...)
	}

	n.node.Flag |= _F_MUT
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

func (n *Node) objSet(key string, val interface{}) error {
	_, t, err := n.objAt(key)
	if err != nil {
		return err
	}
	l := len(n.mut)
	*t = types.Token{
		Kind: types.Type(V_ANY),
		Off: uint32(l),
	}
	n.mut = append(n.mut, val)

	n.node.Flag |= _F_MUT
	return nil
}

func (n *Node) objAdd(key string, val interface{}) error {
	l := len(n.mut)
	k := types.Token{
		Kind: types.T_STRING,
		Flag: _F_KEY,
		Off: uint32(l),
	}
	v := types.Token{
		Kind: types.Type(V_ANY),
		Off: uint32(l+1),
	}
	n.mut = append(n.mut, key)
	n.mut = append(n.mut, val)
	n.node.Kids = append(n.node.Kids, k)
	n.node.Kids = append(n.node.Kids, v)

	n.node.Flag |= _F_MUT
	return nil
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
	if t.Kind == types.Type(V_ANY) {
		x := int(t.Off)
		var right []interface{}
		if x < len(n.mut) - 2 {
			right = n.mut[x+2:]
		}
		n.mut = append(n.mut[:x], right...)
	}

	n.node.Flag |= _F_MUT
	return nil
}

// This will convert a token to Node
//   - scalar type, directly slice original string
//   - array/object, parse to Node for one layer
//   - mut type, use interface{}, which is stored at self.mut[0]
// TODO: handle mut token
func (n *Node) getKidLoad(t types.Token) Node {
	spew.Dump(t)
	if t.Kind != types.Type(V_ANY) {
		return newRawNodeLoad(t.Raw(n.node.JSON), t.Flag)
	} else {
		return NewAny(n.mut[t.Off])
	}
}

func (n *Node) getKidRaw(t types.Token) Node {
	if t.Kind != types.Type(V_ANY) {
		return newRawNode(t.Raw(n.node.JSON), t.Flag)
	} else {
		return NewAny(n.mut[t.Off])
	}
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

func (n *Node) get(key string) Node  {
	if err := n.should(types.T_OBJECT); err != nil {
		return newError(err)
	}
	_, t, err := n.objAt(key)
	if err != nil {
		return newError(err)
	}
	return n.getKidLoad(*t)
}

func (n *Node) index(key int) Node  {
	if err := n.should(types.T_ARRAY); err != nil {
		return newError(err)
	}
	t := n.arrAt(key)
	if t == nil {
		return emptyNode
	}
	return n.getKidLoad(*t)
}

func (n *Node) json(t types.Token) string {
	return t.Raw(n.node.JSON)
}

func (n *Node) str(t types.Token) (string, error) {
	return raw2str(n.json(t), t.Flag.IsEsc(), t.Off)
}

func raw2str(json string, esc bool, off uint32) (string, error) {
	if !esc {
		return json[1: len(json) - 1], nil
	}
	s, err := unquote(json[1: len(json) - 1])
	if err != 0 {
		return "", makeSyntaxError(json, int(off), err.Message())
	} else {
		return s, nil
	}
}

// quoted
func (n *Node) key(t types.Token) (string) {
	if t.Flag & _F_KEY == 0 {
		return n.json(t)
	} else {
		return n.mut[t.Off].(string)
	}
}


// TODO: use flags to make, if is primitives
func parseLazy(json string, path *[]interface{}) (Node, error) {
	// TODO: got real PC of biz caller
	node := Node{}
	node.node.Kids = types.NewToken()

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

	tmp := make([]types.Token, len(node.node.Kids))
	copy(tmp, node.node.Kids)
	types.FreeToken(node.node.Kids)
	node.node.Kids = tmp

	/* check errors */
	if r < 0 {
		return Node{},  makeSyntaxError(json, p, types.ParsingError(-r).Message())
	}

	node.node.JSON = json[r:p]
	types.RecordTokenSize(int64(len(node.node.Kids)))
	return node, nil
}


// Note: not validate the input json, only used internal
func newRawNodeLoad(json string, flag types.Flag) Node {
	n := types.NewNode(json, flag)
	if !n.Kind.IsComplex() {
		return Node{n, nil}
	}
	return NewRaw(json)
}

// Note: not load sub layer, only used for encoding..
func newRawNode(json string, flag types.Flag) Node {
	return Node{types.NewNode(json, flag), nil}
}

