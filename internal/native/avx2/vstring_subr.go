// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__vstring = 48
)

const (
    _stack__vstring = 104
)

const (
    _size__vstring = 2396
)

var (
    _pcsp__vstring = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {2221, 104},
        {2225, 48},
        {2226, 40},
        {2228, 32},
        {2230, 24},
        {2232, 16},
        {2234, 8},
        {2235, 0},
        {2393, 104},
    }
)

var _cfunc_vstring = []loader.CFunc{
    {"_vstring_entry", 0,  _entry__vstring, 0, nil},
    {"_vstring", _entry__vstring, _size__vstring, _stack__vstring, _pcsp__vstring},
}
