package ast

import (
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/rt"
)

func (self *Node) MarshalJSON() ([]byte, error) {
    if err := self.Check(); err != nil {
        return nil, err
    }
    if !self.isMut() {
        return rt.Str2Mem(self.node.JSON), nil
    }
    
}
