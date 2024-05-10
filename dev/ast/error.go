package ast

import (
	"errors"

	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/internal/native/types"
)


var (
    // ErrNotExist means both key and value doesn't exist 
    ErrNotExist error = newError(errors.New("not exist"))

    emptyNode = newError(errors.New("not exist"))

    // ErrUnsupportType means API on the node is unsupported
    ErrUnsupportType error = newError(errors.New("not supported type"))
)

func newError(err error) Node {
    return Node{
        node: types.Node{
            Kind: V_ERROR,
            JSON: err.Error(),
        },
    }
}

// Error returns error message if the node is invalid
func (self Node) Error() string {
    if self.node.Kind != V_ERROR {
        return ""
    } else {
		return self.node.JSON
    } 
}

func makeSyntaxError(json string, p int, msg string) decoder.SyntaxError {
	return decoder.SyntaxError{
		Pos : p,
		Src : json,
		Msg: msg,
	}
}
