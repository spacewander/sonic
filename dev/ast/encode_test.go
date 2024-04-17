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
	 `encoding/json`
	 `testing`
	 `strings`
 
	 `github.com/stretchr/testify/assert`
	 `github.com/stretchr/testify/require`
 )
 
 func TestEncodeValue(t *testing.T) {
	 type Case struct {
		 node Node
		 exp string
		 err bool
	 }
	 input := []Case{
		 {NewNull(), "null", false},
		 {NewBool(true), "true", false},
		 {NewBool(false), "false", false},
		 {NewNumber("0.0"), "0.0", false},
		 {NewString(""), `""`, false},
		 {NewString(`\"\"`), `"\\\"\\\""`, false},
		 {NewArray([]Node{}), "[]", false},
		 {NewArray([]Node{NewString(""), NewNull()}), `["",null]`, false},
		 {NewArray([]Node{NewBool(true), NewString("true"), NewString("\t")}), `[true,"true","\t"]`, false},
		 {NewObject([]Pair{{"a", NewNull()}, {"b", NewNumber("0")}}), `{"a":null,"b":0}`, false},
		 {NewObject([]Pair{{"\ta", NewString("\t")}, {"\bb", NewString("\b")}, {"\nb", NewString("\n")}, {"\ra", NewString("\r")}}),`{"\ta":"\t","\u0008b":"\u0008","\nb":"\n","\ra":"\r"}`, false},
		 {NewObject([]Pair{}), `{}`, false},
		 {NewObject([]Pair{{Key: "", Value: NewNull()}}), `{"":null}`, false},
		 {NewBytes([]byte("hello, world")), `"aGVsbG8sIHdvcmxk"`, false},
		 {NewRaw(`[{ }]`), "[{}]", false},
		 {Node{}, "", true},
		 {NewInt(-1), "-1", false},
		 {NewInt(0), "0", false},
		 {NewUint(0), "0", false},
		 {NewUint(1), "1", false},
		 {NewFloat(1.23), "1.23", false},
	 }
	 for i, c := range input {
		println("test case is ", c.exp)
		 buf, err := json.Marshal(&c.node)
		 if c.err {
			assert.Error(t, err)
			continue
		 }
		 assert.NoError(t, err)
		 if err != nil {
			 t.Fatal(i, err)
		 }
		 assert.Equal(t, c.exp, string(buf), c.exp)
	 }
 }
 
 func TestEncodeNode(t *testing.T) {
	 data := `{"a":[{},[],-0.1,true,false,null,""],"b":0,"c":true,"d":false,"e":null,"g":""}`
	 root, e := NewSearcher(data).GetByPath()
	 if e != nil {
		 t.Fatal(root)
	 }
	 ret, err := root.MarshalJSON()
	 if err != nil {
		 t.Fatal(err)
	 }
	 if string(ret) != data {
		 t.Fatal(string(ret))
	 }
 }

 func TestEncodeNone(t *testing.T) {
	 n := NewObject([]Pair{{Key:"a", Value:Node{}}})
	 out, err := n.MarshalJSON()
	 require.NoError(t, err)
	 require.Equal(t, "{}", string(out))
	 n = NewObject([]Pair{{Key:"a", Value:NewNull()}, {Key:"b", Value:Node{}}})
	 out, err = n.MarshalJSON()
	 require.NoError(t, err)
	 require.Equal(t, `{"a":null}`, string(out))
 
	 n = NewArray([]Node{Node{}})
	 out, err = n.MarshalJSON()
	 require.NoError(t, err)
	 require.Equal(t, "[]", string(out))
	 n = NewArray([]Node{NewNull(), Node{}})
	 out, err = n.MarshalJSON()
	 require.NoError(t, err)
	 require.Equal(t, `[null]`, string(out))
 }
 
 
 type Path = []interface{}
 
 type testGetApi struct {
	 json      string
	 path      Path
 }
 
 type checkError func(error) bool
 
 func isSyntaxError(err error) bool {
	 if err == nil {
		 return false
	 }
	 return strings.HasPrefix(err.Error(), `"Syntax error at index`)
 }
 
 func isEmptySource(err error) bool {
	 if err == nil {
		 return false
	 }
	 return strings.Contains(err.Error(), "no sources available")
 }
 
 func isErrNotExist(err error) bool {
	 return err == ErrNotExist
 }
 
 func isErrUnsupportType(err error) bool {
	 return err == ErrUnsupportType
 }
 
 func testSyntaxJson(t *testing.T, json string, path ...interface{}) {
	 search := NewSearcher(json)
	 _, err := search.GetByPath(path...)
	 assert.True(t, isSyntaxError(err))
 }
 
 func TestGetFromEmptyJson(t *testing.T) {
	 tests := []testGetApi {
		 { "", nil },
		 { "", Path{}},
		 { "", Path{""}},
		 { "", Path{0}},
		 { "", Path{"", ""}},
	 }
	 for _, test := range tests {
		 f := func(t *testing.T) {
			 search := NewSearcher(test.json)
			 _, err := search.GetByPath(test.path...)
			 assert.True(t, isEmptySource(err))
		 }
		 t.Run(test.json, f)
	 }
 }
 
 func TestGetFromSyntaxError(t *testing.T) {
	 tests := []testGetApi {
		 { " \r\n\f\t", Path{} },
		 { "123.", Path{} },
		 { "+124", Path{} },
		 { "-", Path{} },
		 { "-e123", Path{} },
		 { "-1.e123", Path{} },
		 { "-12e456.1", Path{} },
		 { "-12e.1", Path{} },
		 { "[", Path{} },
		 { "{", Path{} },
		 { "[}", Path{} },
		 { "{]", Path{} },
		 { "{,}", Path{} },
		 { "[,]", Path{} },
		 { "tru", Path{} },
		 { "fals", Path{} },
		 { "nul", Path{} },
		 { `{"a":"`, Path{"a"} },
		 { `{"`, Path{} },
		 { `"`, Path{} },
		 { `"\"`, Path{} },
		 { `"\\\"`, Path{} },
		 { `"hello`, Path{} },
		 { `{{}}`, Path{} },
		 { `{[]}`, Path{} },
		 { `{:,}`, Path{} },
		 { `{test:error}`, Path{} },
		 { `{":true}`, Path{} },
		 { `{"" false}`, Path{} },
		 { `{ "" : "false }`, Path{} },
		 { `{"":"",}`, Path{} },
		 { `{ " test : true}`, Path{} },
		 { `{ "test" : tru }`, Path{} },
		 { `{ "test" : true , }`, Path{} },
		 { `{ {"test" : true , } }`, Path{} },
		 { `{"test":1. }`, Path{} },
		 { `{"\\\""`, Path{} },
		 { `{"\\\"":`, Path{} },
		 { `{"\\\":",""}`, Path{} },
		 { `[{]`, Path{} },
		 { `[tru]`, Path{} },
		 { `[-1.]`, Path{} },
		 { `[[]`, Path{} },
		 { `[[],`, Path{} },
		 { `[ true , false , [ ]`, Path{} },
		 { `[true, false, [],`, Path{} },
		 { `[true, false, [],]`, Path{} },
		 { `{"key": [true, false, []], "key2": {{}}`, Path{} },
	 }
 
	 for _, test := range tests {
		 f := func(t *testing.T) {
			 testSyntaxJson(t, test.json, test.path...)
			 path := append(Path{"key"}, test.path...)
			 testSyntaxJson(t, `{"key":` + test.json, path...)
			 path  = append(Path{""}, test.path...)
			 testSyntaxJson(t, `{"":` + test.json, path...)
			 path  = append(Path{1}, test.path...)
			 testSyntaxJson(t, `["",` + test.json, path...)
		 }
		 t.Run(test.json, f)
	 }
 }
 
