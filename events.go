package idle

import (
	"github.com/faiface/pixel"

	"github.com/faiface/pixel/pixelgl"
)

type Event struct {
	Action   MouseAction
	MousePos pixel.Vec
	Button   pixelgl.Button
	Scroll   pixel.Vec
	Doomed   bool
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

func (action MouseAction) The(button pixelgl.Button, at pixel.Vec) Event {
	return Event{Action: action, Button: button, MousePos: at}
}
