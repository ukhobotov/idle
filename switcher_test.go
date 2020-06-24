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
		Elements: []Element{
			ctr,
		},
	}
	switcher.FitInto(4, 4, 100, 100)
	assert.Equal(t, 10.0, ctr.x1)
	assert.Equal(t, 17.0, ctr.y1)
	assert.Equal(t, 22.0, ctr.x2)
	assert.Equal(t, 48.0, ctr.y2)
}

func TestSwitcher_Handle(t *testing.T) {
	defer func(t *testing.T) {
		assert.Nil(t, recover())
	}(t)
	(&Switcher{Elements: []Element{&Container{}}}).Handle(Event{}, 0, 0)
}

func TestSwitcher_Draw(t *testing.T) {
	defer func(t *testing.T) {
		assert.Nil(t, recover())
	}(t)
	(&Switcher{Elements: []Element{&Container{}}}).Draw(nil)
}
