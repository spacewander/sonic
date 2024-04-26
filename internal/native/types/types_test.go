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

package types

import "testing"

func TestPredictTokenSize(t *testing.T) {
	tests := []struct {
		name string
		last []int64
		lack []bool
		final int64
	}{
		{"default", []int64{8, 8, 8}, []bool{false, false, false}, _DefaultTokenSize*2},
		{"enlarge", []int64{32, 32, 64}, []bool{true, false, true}, 100},
		{"shrink", []int64{128, 1, 1, 1, 1, 1}, []bool{true, false, false, false, false, false, false}, 34},
		{"wave", []int64{256, 1, 1, 1, 1, 1, 1, 32, 1, 1, 1, 256, 1}, []bool{true, false, false, false, false, false, false, false, false, false, false, true, false, false}, 66},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i:=0; i<len(tt.last); i++ {
				got := PredictTokenSize()
				last := tt.last[i]
				println(i, got, last)
				if got < last && !tt.lack[i] || got >= last && tt.lack[i] {
					t.Fatalf("%d time failed, got %d", i, got)
				} 
				RecordTokenSize(last)
			}
			final := PredictTokenSize()
			if final != tt.final {
				t.Errorf("final size = %v, want %v", final, tt.final)
			}
		})
	}
}

func BenchmarkPredictTokenSize(b *testing.B) {
	for i:=0; i<b.N; i++ {
		got := PredictTokenSize()
		RecordTokenSize(got+1)
	}
}