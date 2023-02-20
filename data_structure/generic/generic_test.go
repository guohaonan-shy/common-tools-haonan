package generic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ContainInt64(t *testing.T) {
	res := Contains[int64](int64(1), []int64{1, 2, 3})
	res2 := Contains[int64](int64(1123121), []int64{1, 2, 3})
	assert.Equal(t, true, res)
	assert.Equal(t, false, res2)
}

func Test_ContainString(t *testing.T) {
	res := Contains("abc", []string{"abc", "def", "efs"})
	res2 := Contains("a", []string{"b", "c", "d"})
	assert.Equal(t, true, res)
	assert.Equal(t, false, res2)
}
