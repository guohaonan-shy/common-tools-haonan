package every_day

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_650(t *testing.T) {
	t.Run("650", func(t *testing.T) {
		assert.Equal(t, 6, minSteps(9))
	})
}
