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
	n1, err := NewParser(`[1,"1",true,null]`).Parse()
	require.NoError(t, err)
	require.Equal(t, len(n1.Kids), 4)

	n1, err = NewParser(`[]`).Parse()
	require.NoError(t, err)
	require.Equal(t, len(n1.Kids), 0)

	n1, err = NewParser(`{}`).Parse()
	require.NoError(t, err)
	require.Equal(t, len(n1.Kids), 0)

	n1, err = NewParser(`{"key": null, "k2": {}}`).Parse()
	require.NoError(t, err)
	spew.Dump(n1.Kids, len(n1.Kids))
	require.Equal(t, len(n1.Kids), 4)

	src := getSample(100, 0)
	n, err := NewParser(src).Parse()
	require.NoError(t, err)
	n50 := n.GetByPath("50")
	require.Empty(t, n50.Error())
	v, _ := n50.Int64()
	require.Equal(t, int64(1), v)
	js, err := n.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, src, string(js))
	src = getSample(100, 1)
	n, err = NewParser(src).Parse()
	require.NoError(t, err)
	js, err = n.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, src, string(js))
}

func BenchmarkNode_GetByPath(b *testing.B) {
	b.Run("10/2", func(b *testing.B) {
		src := getSample(10, 0)
		b.ResetTimer()
		n, _ := NewParser(src).Parse()
		for i:=0; i< b.N; i++ {
			_ = n.GetByPath("5")
		}
	})
	b.Run("100/2", func(b *testing.B) {
		src := getSample(100, 0)
		b.ResetTimer()
		n, _ := NewParser(src).Parse()
		for i:=0; i< b.N; i++ {
			_ = n.GetByPath("50")
		}
	})
	b.Run("1000/2", func(b *testing.B) {
		src := getSample(1000, 0)
		b.ResetTimer()
		n, _ := NewParser(src).Parse()
		for i:=0; i< b.N; i++ {
			_ = n.GetByPath("500")
		}
	})
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

