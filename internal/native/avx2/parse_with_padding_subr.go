// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__parse_with_padding = 336
)

const (
    _stack__parse_with_padding = 192
)

const (
    _size__parse_with_padding = 47668
)

var (
    _pcsp__parse_with_padding = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {6358, 192},
        {6365, 48},
        {6366, 40},
        {6368, 32},
        {6370, 24},
        {6372, 16},
        {6374, 8},
        {6375, 0},
        {47668, 192},
    }
)

var _cfunc_parse_with_padding = []loader.CFunc{
    {"_parse_with_padding_entry", 0,  _entry__parse_with_padding, 0, nil},
    {"_parse_with_padding", _entry__parse_with_padding, _size__parse_with_padding, _stack__parse_with_padding, _pcsp__parse_with_padding},
}
