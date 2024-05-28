// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__quote = 48
)

const (
    _stack__quote = 80
)

const (
    _size__quote = 1760
)

var (
    _pcsp__quote = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1701, 80},
        {1705, 48},
        {1706, 40},
        {1708, 32},
        {1710, 24},
        {1712, 16},
        {1714, 8},
        {1715, 0},
        {1750, 80},
    }
)

var _cfunc_quote = []loader.CFunc{
    {"_quote_entry", 0,  _entry__quote, 0, nil},
    {"_quote", _entry__quote, _size__quote, _stack__quote, _pcsp__quote},
}
