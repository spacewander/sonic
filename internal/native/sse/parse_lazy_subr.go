// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__parse_lazy = 12480
    _entry__load_lazy = 240
)

const (
    _stack__parse_lazy = 1000
    _stack__load_lazy = 784
)

const (
    _size__parse_lazy = 12768
    _size__load_lazy = 8520
)

var (
    _pcsp__parse_lazy = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {12475, 216},
        {12482, 48},
        {12483, 40},
        {12485, 32},
        {12487, 24},
        {12489, 16},
        {12491, 8},
        {12492, 0},
        {12587, 216},
        {12594, 48},
        {12595, 40},
        {12597, 32},
        {12599, 24},
        {12601, 16},
        {12603, 8},
        {12604, 0},
        {12768, 216},
    }
    _pcsp__load_lazy = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {7956, 216},
        {7963, 48},
        {7964, 40},
        {7966, 32},
        {7968, 24},
        {7970, 16},
        {7972, 8},
        {7973, 0},
        {8520, 216},
    }
)

var _cfunc_parse_lazy = []loader.CFunc{
    {"_parse_lazy_entry", 0,  _entry__parse_lazy, 0, nil},
    {"_parse_lazy", _entry__parse_lazy, _size__parse_lazy, _stack__parse_lazy, _pcsp__parse_lazy},
}
var _cfunc_load_lazy = []loader.CFunc{
    {"_load_lazy_entry", 0,  _entry__load_lazy, 0, nil},
    {"_load_lazy", _entry__load_lazy, _size__load_lazy, _stack__load_lazy, _pcsp__load_lazy},
}
