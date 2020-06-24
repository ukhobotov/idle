package carbon

type Switcher struct {
	Location
	Style   *Style
	Content []Element
	Current Element
}

func (switcher *Switcher) FitInto(x1, y1, x2, y2 float64) {
	switcher.Location.FitInto(x1, y1, x2, y2)
	for _, element := range switcher.Content {
		element.FitInto(switcher.Location.Absolute())
	}
}

func (switcher *Switcher) Rasterize() {
	switcher.Style.Rasterize(switcher.Size())
	for _, element := range switcher.Content {
		element.Rasterize()
	}
}

func (switcher *Switcher) Handle(event Event, x, y float64) {
	switcher.Current.Handle(event, x, y)
}

func (switcher *Switcher) Draw(win *Window) {
	x, y := switcher.Center()
	switcher.Style.Draw(win, x, y)
	if switcher.Content != nil {
		switcher.Current.Draw(win)
	}
}
