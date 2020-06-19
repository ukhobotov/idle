package carbon

import (
	"github.com/faiface/pixel"

	"github.com/faiface/pixel/pixelgl"
)

type Event struct {
	Action MouseAction
	Button pixelgl.Button
	Scroll pixel.Vec
	Doomed bool
}

var MoveEvent = Event{Action: Move, Button: NilButton}

type MouseAction int

const (
	Move MouseAction = iota
	Press
	Release
	Scroll

	NilButton pixelgl.Button = -1
)

func (action MouseAction) The(button pixelgl.Button) Event {
	return Event{Action: action, Button: button}
}
