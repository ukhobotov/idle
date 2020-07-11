package idle

import "github.com/faiface/pixel"

type Container struct {
	Location
	State   *State
	Content []Element
}

func (ctr *Container) FitInto(x1, y1, x2, y2 float64) {
	ctr.Location.FitInto(x1, y1, x2, y2)
	for _, element := range ctr.Content {
		element.FitInto(ctr.Absolute())
	}
}

func (ctr *Container) Handle(event Event) {
	for _, element := range ctr.Content {
		if element.Contains(event.MousePos.XY()) {
			element.Handle(event)
		}
		if event.Doomed {
			return
		}
	}
}

func (ctr *Container) Rasterize() {
	ctr.State.Rasterize(ctr.Location.Size())
	for _, element := range ctr.Content {
		element.Rasterize()
	}
}

func (ctr *Container) Draw(target pixel.Target) {
	ctr.State.Draw(target, ctr.Center())
	for _, element := range ctr.Content {
		element.Draw(target)
	}
}
