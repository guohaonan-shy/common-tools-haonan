package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Stack(t *testing.T) {
	t.Run("postfix_expression", func(t *testing.T) {
		input := []string{"2", "3", "+", "4", "*"}
		result := PostfixExpression(input)
		assert.Equal(t, 20, result)
	})
}
