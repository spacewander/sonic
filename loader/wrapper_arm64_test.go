/**
* Copyright 2023 ByteDance Inc.
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

package loader

import (
	`os`
	`runtime`
	`runtime/debug`
	`testing`
	`time`

	`github.com/stretchr/testify/require`
)

var (
	debugAsyncGC = os.Getenv("SONIC_NO_ASYNC_GC") == ""
)

func TestMain(m *testing.M) {
	go func ()  {
		if !debugAsyncGC {
			return
		}
		println("Begin GC looping...")
		for {
		runtime.GC()
		debug.FreeOSMemory() 
		}
		println("stop GC looping!")
	}()
	time.Sleep(time.Millisecond*100)
	m.Run()
}

func TestWrapC(t *testing.T) {
	var stub func(a int64, val *int64) (ret int64) 
	/**
	int64_t func (int64_t a, int64_t *b) {
		return a + *b;
	}
	**/
	ct := []byte{
		0xfd, 0x7b, 0xbf, 0xa9, // stp	x29, x30, [sp, #-16]!
    	0x21, 0x00, 0x40, 0xf9, // ldr x1, [x1]
    	0x20, 0x00, 0x00, 0x8b, // add x0, x1, x0
		0xfd, 0x7b, 0xc1, 0xa8, // ldp	x29, x30, [sp], #16
    	0xc0, 0x03, 0x5f, 0xd6, // ret
	}
	size := uint32(len(ct))
	WrapGoC(ct, []CFunc{{
		Name:     "add",
		EntryOff: 0,
		TextSize: size,
		MaxStack: uintptr(40),
		Pcsp:     [][2]uint32{
			{1, 0},
			{size-4, 24},
			{size, 0},
		},
	}}, []GoC{{
		CName:     "add",
		GoFunc:   &stub,
	} }, "dummy/native", "dummy/native.c")

	f := stub
	b := int64(2)    
	println("b : ", &b)
	var c *int64 = &b
	runtime.SetFinalizer(c, func(x *int64){
		println("c got GC: ", x)
	})
	// runtime.GC()
	println("before")
	var act int64
	testFunc(func() {
		act = f(1, c)
	})
	println("after")
	// runtime.GC()
	require.Equal(t, int64(3), act)
}
