package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitcher_FitInto(t *testing.T) {
	ctr := &Container{}
	switcher := &Switcher{
		Location: Location{
			Left:   Margin(6),
			Bottom: Margin(13),
			Right:  Margin(18),
			Top:    Margin(-52),
		},
		Content: []Element{
			ctr,
		},
	}
	switcher.FitInto(4, 4, 100, 100)
	assert.Equal(t, 10.0, ctr.x1)
	assert.Equal(t, 17.0, ctr.y1)
	assert.Equal(t, 22.0, ctr.x2)
	assert.Equal(t, 48.0, ctr.y2)
}
