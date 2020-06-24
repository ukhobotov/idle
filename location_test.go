package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstraint_Absolute(t *testing.T) {
	assert.Equal(t, 6.0, Constraint{}.Absolute(6, 17, false))
	assert.Equal(t, 19.0, Constraint{}.Absolute(4, 19, true))
	assert.Equal(t, 22.0, Constraint{Margin: 18}.Absolute(4, 100, true))
}
