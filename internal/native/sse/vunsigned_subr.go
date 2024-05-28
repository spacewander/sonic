// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__vunsigned = 0
)

const (
    _stack__vunsigned = 24
)

const (
    _size__vunsigned = 356
)

var (
    _pcsp__vunsigned = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {72, 24},
        {73, 16},
        {75, 8},
        {76, 0},
        {87, 24},
        {88, 16},
        {90, 8},
        {91, 0},
        {114, 24},
        {115, 16},
        {117, 8},
        {118, 0},
        {281, 24},
        {282, 16},
        {284, 8},
        {285, 0},
        {336, 24},
        {337, 16},
        {339, 8},
        {340, 0},
        {348, 24},
        {349, 16},
        {351, 8},
        {353, 0},
    }
)

var _cfunc_vunsigned = []loader.CFunc{
    {"_vunsigned_entry", 0,  _entry__vunsigned, 0, nil},
    {"_vunsigned", _entry__vunsigned, _size__vunsigned, _stack__vunsigned, _pcsp__vunsigned},
}
