
package ast

import (
    `github.com/bytedance/sonic/internal/rt`
)

func (self *Node) MarshalJSON() ([]byte, error) {
    if !self.IsMut() {
        return rt.Str2Mem(self.JSON), nil
    }
    // hanlde mutates
    panic("not implement!")
}
