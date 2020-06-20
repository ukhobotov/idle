package components

import "carbon"

type Button struct {
	carbon.Location
	carbon.Handler

	Text carbon.Text
}

func (button *Button) Rasterize() {
	// panic("implement me")
}

func (button *Button) Handle(event carbon.Event, x, y float64) {
	// panic("implement me")
}

func (button *Button) Draw(*carbon.Window) {
	// panic("implement me")
}
