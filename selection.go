package idle

type Hovered interface {
	Contains(x, y float64) bool
	Dehover()
}

var hovered []Hovered

func Hover(hover Hovered) {
	hovered = append(hovered, hover)
}

func HandleHovered(event Event) {
	if event.Action == Move && event.Button != NilButton {
		return
	}
	for i, hover := range hovered {
		if !hover.Contains(event.MousePos.XY()) {
			hover.Dehover()
			last := len(hovered) - 1
			hovered[i], hovered[last] = hovered[last], nil
			hovered = hovered[:last]
		}
	}
}

type Focused interface {
	Defocus()
}

var focused Focused

func Focus(focus Focused) {
	if focused != nil {
		focused.Defocus()
	}
	focused = focus
}
