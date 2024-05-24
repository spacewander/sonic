// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__get_by_path = 224
)

const (
    _stack__get_by_path = 272
)

const (
    _size__get_by_path = 20912
)

var (
    _pcsp__get_by_path = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {20158, 272},
        {20159, 264},
        {20161, 256},
        {20163, 248},
        {20165, 240},
        {20167, 232},
        {20171, 224},
        {20912, 272},
    }
)

var _cfunc_get_by_path = []loader.CFunc{
    {"_get_by_path_entry", 0,  _entry__get_by_path, 0, nil},
    {"_get_by_path", _entry__get_by_path, _size__get_by_path, _stack__get_by_path, _pcsp__get_by_path},
}
