// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
`github.com/bytedance/sonic/internal/rt`
)

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __f64toa_entry() uintptr

var (
    _subr__f64toa uintptr = rt.GetFuncPC(__f64toa_entry) + 46
)

const (
    _stack__f64toa = 56
)

var (
    _ = _subr__f64toa
)

const (
    _ = _stack__f64toa
)
