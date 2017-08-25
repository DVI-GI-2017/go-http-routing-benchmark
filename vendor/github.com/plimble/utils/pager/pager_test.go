package pager

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPage(t *testing.T) {
	assert := assert.New(t)

	c, p := Page(5, 0, 10)
	assert.Equal(1, c)
	assert.Equal(2, p)
	c, p = Page(10, 21, 30)
	assert.Equal(3, c)
	assert.Equal(3, p)
	c, p = Page(10, 21, 0)
	assert.Equal(0, c)
	assert.Equal(0, p)
	c, p = Page(10, 0, 1)
	assert.Equal(1, c)
	assert.Equal(1, p)
	c, p = Page(10, 5, 0)
	assert.Equal(0, c)
	assert.Equal(0, p)
	c, p = Page(50, 200, 100)
	assert.Equal(0, c)
	assert.Equal(0, p)
}

func TestPageToRange(t *testing.T) {
	o := Offset(3, 2)
	assert.Equal(t, 3, o)
	o = Offset(10, 1)
	assert.Equal(t, 0, o)
	o = Offset(10, 2)
	assert.Equal(t, 10, o)
}

func BenchmarkGetPage(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Page(5, 6, 10)
	}
}
