package ast

import (
	"fmt"

	"github.com/bytedance/sonic/internal/encoder/alg"
	"github.com/bytedance/sonic/internal/native/types"
	uq "github.com/bytedance/sonic/unquote"
	"github.com/cloudwego/base64x"
)

// NOTICE: must assume pos is at ']' or '}'
func backward(json string, pos int) (stop int, empty bool) {
    if json[stop] != ']' && json[stop] != '}' {
        panic(fmt.Sprintf("char '%c' at stop %d is not close dilimeter of JSON", json[stop], stop))
    }
    stop = pos - 1
    for ; stop >= 0 && isSpace(json[stop]); stop-- {
    }
    empty = (json[stop] == '[' || json[stop] == '{')
    stop += 1
    return
}

func decodeBase64(src string) ([]byte, error) {
    return base64x.StdEncoding.DecodeString(src)
}

func encodeBase64(src []byte) string {
    return base64x.StdEncoding.EncodeToString(src)
}

func unquote(src string) (string, types.ParsingError) {
    return uq.String(src)
}

// [2,"a"],1 => {"a":1}
// ["a",2],1  => "a":[1]
func makePathAndValue(b []byte, path []interface{}, allowAppend bool, val Node) ([]byte, error) {
	for i, k := range path {
		if key, ok := k.(string); ok {
			b = alg.Quote(b, key, false)
			b = append(b, ":"...)
		}
		if i == len(path)-1 {
			b = append(b, val.node.JSON...)
			break
		}
		n := path[i+1]
		if _, ok := n.(int); ok {
			if !allowAppend {
				return nil, ErrNotExist
			}
			b = append(b, "["...)
		} else if _, ok := n.(string); ok {
			b = append(b, `{`...)
		} else {
			panic("path must be either int or string")
		}
	}
	for i := len(path) - 1; i >= 1; i-- {
		k := path[i]
		if _, ok := k.(int); ok {
			b = append(b, "]"...)
		} else if _, ok := k.(string); ok {
			b = append(b, `}`...)
		}
	}
	return b, nil
}
