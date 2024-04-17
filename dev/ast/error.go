package ast

import (
    `fmt`
    `strings`
    `unsafe`
	`errors`
)


var (
    // ErrNotExist means both key and value doesn't exist 
    ErrNotExist error = newError(errors.New("not exist"))

    // ErrUnsupportType means API on the node is unsupported
    ErrUnsupportType error = newError(errors.New("not exist"))
)

func newError(err error) *Node {
    return &Node{
        typ: _V_ERROR,
        ptr: unsafe.Pointer(&err),
    }
}

// Error returns error message if the node is invalid
func (self Node) Error() string {
    if self.typ != _V_ERROR {
        return ""
    } else {
		err := *(*error)(self.ptr)
        return err.Error()
    } 
}


func newSyntaxError(err SyntaxError) *Node {
	var e error = err
    return &Node{
        typ: _V_ERROR,
        ptr: unsafe.Pointer(&e),
    }
}

type SyntaxError struct {
    Pos  int
    Src  string
    Msg  string
}

func (self SyntaxError) Error() string {
    return fmt.Sprintf("%q", self.Description())
}

func (self SyntaxError) Description() string {
    return "Syntax error " + self.description()
}

func (self SyntaxError) description() string {
    i := 16
    p := self.Pos - i
    q := self.Pos + i

    /* check for empty source */
    if self.Src == "" {
        return fmt.Sprintf("no sources available: %#v", self)
    }

    /* prevent slicing before the beginning */
    if p < 0 {
        p, q, i = 0, q - p, i + p
    }

    /* prevent slicing beyond the end */
    if n := len(self.Src); q > n {
        n = q - n
        q = len(self.Src)

        /* move the left bound if possible */
        if p > n {
            i += n
            p -= n
        }
    }

    /* left and right length */
    x := clamp_zero(i)
    y := clamp_zero(q - p - i - 1)

    /* compose the error description */
    return fmt.Sprintf(
        "at index %d: %s\n\n\t%s\n\t%s^%s\n",
        self.Pos,
        self.Message(),
        self.Src[p:q],
        strings.Repeat(".", x),
        strings.Repeat(".", y),
    )
}

func (self SyntaxError) Message() string {
    return self.Msg
}

func clamp_zero(v int) int {
    if v < 0 {
        return 0
    } else {
        return v
    }
}

