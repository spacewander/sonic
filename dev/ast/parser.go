/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import (
	`github.com/bytedance/sonic/internal/native`
	`github.com/bytedance/sonic/internal/native/types`
	`runtime`
)

type Parser struct {
    pos	int
    src	string
}

/** Parser Interface **/

func (self *Parser) Pos() int {
    return self.pos
}

func (self *Parser) Parse() (Node, error) {
	node, err := parseLazy(self.src, nil)
	if err != nil {
		return Node{}, err
	}
	return node, nil
}

// NewParser returns pointer of new allocated parser
func NewParser(src string) *Parser {
    return &Parser{src: src}
}

func (self *Parser) getByPath(path ...interface{}) (Node, error) {
    return parseLazy(self.src[self.pos:], &path)
}

func (self *Parser) skip(path ...interface{}) (int, error) {
    fsm := types.NewStateMachine()
    start := native.GetByPath(&self.src, &self.pos, &path, fsm)
    types.FreeStateMachine(fsm)
    runtime.KeepAlive(path)
    if start < 0 {
        if -start == int(types.ERR_NOT_FOUND) {
            return -1, ErrNotExist
        }
        return self.pos, makeSyntaxError(self.src, self.pos, types.ParsingError(-start).Message())
    }
    return start, nil
}


