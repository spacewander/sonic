// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
`github.com/bytedance/sonic/internal/rt`
)

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __validate_utf8_entry() uintptr

var (
    _subr__validate_utf8 uintptr = rt.GetFuncPC(__validate_utf8_entry) + 16
)

const (
    _stack__validate_utf8 = 48
)

var (
    _ = _subr__validate_utf8
)

const (
    _ = _stack__validate_utf8
)
