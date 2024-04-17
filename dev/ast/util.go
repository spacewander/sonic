package ast

import (
	`github.com/cloudwego/base64x`
    `github.com/bytedance/sonic/internal/native/types`
    uq `github.com/bytedance/sonic/unquote`
)


func decodeBase64(src string) ([]byte, error) {
    return base64x.StdEncoding.DecodeString(src)
}

func encodeBase64(src []byte) string {
    return base64x.StdEncoding.EncodeToString(src)
}

func unquote(src string) (string, types.ParsingError) {
    return uq.String(src)
}

