package ast

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func getSample(width int, depth int) string {
	obj := map[string]interface{}{}
	for i := 0; i < width; i++ {
		var v  interface{}
		if depth > 0 {
			v = json.RawMessage(getSample(width/2+1, depth-1))
		} else {
			v = 1
		}
		obj[strconv.Itoa(i)] = v
	}
	js, _ := json.Marshal(obj)
	return string(js)
}

func TestNodeParse(t *testing.T) {
	n1, err := NewParser(`[1,"1",true]`).Parse()
	require.NoError(t, err)
	spew.Dump(n1.Kids)
	src := getSample(100, 0)
	n, err := NewParser(src).Parse()
	require.NoError(t, err)
	js, err := n.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, src, string(js))
	src = getSample(100, 1)
	n, err = NewParser(src).Parse()
	spew.Dump(n.Kids)
	require.NoError(t, err)
	js, err = n.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, src, string(js))
}

func BenchmarkParse(b *testing.B) {
	b.Run("10-0", func(b *testing.B) {
		src := getSample(10, 0)
		b.ResetTimer()
		for i:=0; i< b.N; i++ {
			_, _ = NewParser(src).Parse()
		}
	})
	b.Run("10-1", func(b *testing.B) {
		src := getSample(10, 1)
		b.ResetTimer()
		for i:=0; i< b.N; i++ {
			_, _ = NewParser(src).Parse()
		}
	})
	b.Run("100-0", func(b *testing.B) {
		src := getSample(100, 0)
		b.ResetTimer()
		for i:=0; i< b.N; i++ {
			_, _ = NewParser(src).Parse()
		}
	})
	b.Run("100-1", func(b *testing.B) {
		src := getSample(100, 1)
		b.ResetTimer()
		for i:=0; i< b.N; i++ {
			_, _ = NewParser(src).Parse()
		}
	})
	b.Run("1000-0", func(b *testing.B) {
		src := getSample(1000, 0)
		b.ResetTimer()
		for i:=0; i< b.N; i++ {
			_, _ = NewParser(src).Parse()
		}
	})
	b.Run("1000-1", func(b *testing.B) {
		src := getSample(1000, 1)
		b.ResetTimer()
		for i:=0; i< b.N; i++ {
			_, _ = NewParser(src).Parse()
		}
	})
}

