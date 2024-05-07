package ast

import (
	"errors"

	"github.com/bytedance/sonic/internal/native/types"
)


var (
    // ErrNotExist means both key and value doesn't exist 
    ErrNotExist error = newError(errors.New("not exist"))

    emptyNode = newError(errors.New("not exist"))

    // ErrUnsupportType means API on the node is unsupported
    ErrUnsupportType error = newError(errors.New("not supported type"))

    ErrOutOfRange error = newError(errors.New("index out of range"))
)

func newError(err error) Node {
    return Node{
        Node: types.Node{
            Kind: _V_ERROR,
            JSON: err.Error(),
        },
    }
}

// Error returns error message if the node is invalid
func (self Node) Error() string {
    if self.Kind != _V_ERROR {
        return ""
    } else {
		return self.JSON
    } 
}
