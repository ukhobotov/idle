package idle

import "github.com/faiface/pixel"

type Switcher struct {
	Location
	Style    *State
	Elements []Element
	Current  Element
}

func (switcher *Switcher) FitInto(x1, y1, x2, y2 float64) {
	switcher.Location.FitInto(x1, y1, x2, y2)
	for _, element := range switcher.Elements {
		element.FitInto(switcher.Location.Absolute())
	}
}

func (switcher *Switcher) Rasterize() {
	switcher.Style.Rasterize(switcher.Size())
	for _, element := range switcher.Elements {
		element.Rasterize()
	}
}

func (switcher *Switcher) Handle(event Event) {
	if switcher.Current != nil {
		switcher.Current.Handle(event)
	}
}

func (switcher *Switcher) Draw(target pixel.Target) {
	switcher.Style.Draw(target, switcher.Center())
	if switcher.Current != nil {
		switcher.Current.Draw(target)
	}
}
