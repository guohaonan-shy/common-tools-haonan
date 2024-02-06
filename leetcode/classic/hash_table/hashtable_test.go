package hash_table

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_128(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := []int{100, 4, 200, 1, 3, 2}
		assert.Equal(t, 4, longestConsecutive(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
		assert.Equal(t, 9, longestConsecutive(case2))
	})

}

func Test_BloomFilter(t *testing.T) {
	println(BloomHash("sin_6017"))
	fmt.Printf("%x\n", BloomHash("sin_6017"))

	println(BloomFilter("sin_6017"))

	fmt.Printf("%x\n", InsertInBloomFilter("sin_6017"))
}
