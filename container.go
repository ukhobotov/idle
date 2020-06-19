package carbon

type Container struct {
	Layout
	Style   *Style
	Handler *Handler
	Content []Element
}

func (ctr *Container) FitInto(x1, y1, x2, y2 float64) {
	ctr.Layout.FitInto(x1, y1, x2, y2)
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
		ctr.Style.Rasterize(ctr.Layout.Size())
	}
	for _, element := range ctr.Content {
		element.Rasterize()
	}
}

func (ctr *Container) Draw() {
	if ctr.Style != nil {
		ctr.Style.Draw(ctr.Layout.Center())
	}
	for _, element := range ctr.Content {
		element.Draw()
	}
}
