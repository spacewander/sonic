// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__skip_number = 112
)

const (
    _stack__skip_number = 72
)

const (
    _size__skip_number = 1060
)

var (
    _pcsp__skip_number = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {879, 72},
        {883, 48},
        {884, 40},
        {886, 32},
        {888, 24},
        {890, 16},
        {892, 8},
        {893, 0},
        {1060, 72},
    }
)

var _cfunc_skip_number = []loader.CFunc{
    {"_skip_number_entry", 0,  _entry__skip_number, 0, nil},
    {"_skip_number", _entry__skip_number, _size__skip_number, _stack__skip_number, _pcsp__skip_number},
}
