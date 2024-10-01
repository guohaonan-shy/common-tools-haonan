package every_day

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_650(t *testing.T) {
	t.Run("650", func(t *testing.T) {
		assert.Equal(t, 0, minSteps(1))
		assert.Equal(t, 6, minSteps(9))
		assert.Equal(t, 3, minSteps(3))
	})
}

func Test_688(t *testing.T) {
	t.Run("knight", func(t *testing.T) {
		assert.Equal(t, 0.06250, knightProbability(3, 2, 0, 0))
	})
}

func Test_983(t *testing.T) {
	t.Run("minimum travel ticket case1", func(t *testing.T) {
		assert.Equal(t, 11, minimumTravelCost([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	})

	t.Run("minimum travel ticket case2", func(t *testing.T) {
		assert.Equal(t, 17, minimumTravelCost([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
	})

	t.Run("minimum travel ticket case2", func(t *testing.T) {
		assert.Equal(t, 6, minimumTravelCost([]int{1, 4, 6, 7, 8, 20}, []int{7, 2, 15}))
	})

	t.Run("minimum travel ticket case3", func(t *testing.T) {
		assert.Equal(t, 50, minimumTravelCost([]int{1, 2, 3, 4, 6, 8, 9, 10, 13, 14, 16, 17, 19, 21, 24, 26, 27, 28, 29}, []int{3, 14, 50}))
	})
}

func Test_1570(t *testing.T) {
	t.Run("dot product", func(t *testing.T) {
		v1 := Constructor([]int{1, 0, 0, 2, 3})
		v2 := Constructor([]int{0, 3, 0, 4, 0})

		assert.Equal(t, 8, v1.dotProduct(v2))
	})
}

//func Test_1531(t *testing.T) {
//	t.Run("compressed string", func(t *testing.T) {
//		assert.Equal(t, 4, getLengthOfOptimalCompression("aaabcccd", 2))
//	})
//}
