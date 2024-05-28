// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__i64toa = 80
)

const (
    _stack__i64toa = 8
)

const (
    _size__i64toa = 2320
)

var (
    _pcsp__i64toa = [][2]uint32{
        {1, 0},
        {173, 8},
        {174, 0},
        {512, 8},
        {513, 0},
        {646, 8},
        {647, 0},
        {1123, 8},
        {1124, 0},
        {1263, 8},
        {1264, 0},
        {1579, 8},
        {1580, 0},
        {1942, 8},
        {1943, 0},
        {2312, 8},
        {2314, 0},
    }
)

var _cfunc_i64toa = []loader.CFunc{
    {"_i64toa_entry", 0,  _entry__i64toa, 0, nil},
    {"_i64toa", _entry__i64toa, _size__i64toa, _stack__i64toa, _pcsp__i64toa},
}
