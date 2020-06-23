package carbon

type Container struct {
	Bounds
	Style   *Style
	Content []Element
}

func (ctr *Container) FitInto(x1, y1, x2, y2 float64) {
	ctr.Bounds.FitInto(x1, y1, x2, y2)
	for _, element := range ctr.Content {
		element.FitInto(ctr.Absolute())
	}
}

func (ctr *Container) Handle(event Event, x, y float64) {
	for _, element := range ctr.Content {
		if element.Contains(x, y) {
			element.Handle(event, x, y)
		}
		if event.Doomed {
			return
		}
	}
}

func (ctr *Container) Rasterize() {
	if ctr.Style != nil {
		ctr.Style.Rasterize(ctr.Bounds.Size())
	}
	for _, element := range ctr.Content {
		element.Rasterize()
	}
}

func (ctr *Container) Draw(win *Window) {
	if ctr.Style != nil {
		x, y := ctr.Bounds.Center()
		ctr.Style.Draw(win, x, y)
	}
	for _, element := range ctr.Content {
		element.Draw(win)
	}
}
