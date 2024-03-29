package every_day

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
