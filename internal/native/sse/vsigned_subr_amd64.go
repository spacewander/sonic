// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
`github.com/bytedance/sonic/internal/rt`
)

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __vsigned_entry() uintptr

var (
    _subr__vsigned uintptr = rt.GetFuncPC(__vsigned_entry) + 16
)

const (
    _stack__vsigned = 16
)

var (
    _ = _subr__vsigned
)

const (
    _ = _stack__vsigned
)
