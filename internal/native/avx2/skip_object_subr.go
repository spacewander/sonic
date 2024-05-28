// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__skip_object = 176
)

const (
    _stack__skip_object = 160
)

const (
    _size__skip_object = 10428
)

var (
    _pcsp__skip_object = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {9907, 160},
        {9911, 48},
        {9912, 40},
        {9914, 32},
        {9916, 24},
        {9918, 16},
        {9920, 8},
        {9921, 0},
        {10428, 160},
    }
)

var _cfunc_skip_object = []loader.CFunc{
    {"_skip_object_entry", 0,  _entry__skip_object, 0, nil},
    {"_skip_object", _entry__skip_object, _size__skip_object, _stack__skip_object, _pcsp__skip_object},
}
