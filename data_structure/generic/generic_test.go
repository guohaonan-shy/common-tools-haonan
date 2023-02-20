package generic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ContainInt64(t *testing.T) {
	res := Contains[int64](int64(1), []int64{1, 2, 3})
	assert.Equal(t, true, res)
}
