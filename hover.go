package idle

import (
	"github.com/faiface/pixel"
)

type (
	// Handler is a basic element for mouse action handling as press, release or hover.
	// It also can be styled for different states using HandlerStyle. Most common are stored in styles package.
	Handler struct {
		Location
		Style    HandlerStyle
		Final    DrawerFunc
		Disabled bool

		OnPress, OnRelease, OnHover func(this *Handler)

		background, foreground     *State
		hovered, pressed, selected bool
	}

	// HandlerStyle is a set of styles for the Handler.
	HandlerStyle struct {
		Idle, Hover, Active, Focus, Disabled *State

		Common DrawerFunc
	}
)

func (handler *Handler) Handle(event Event) {
	if handler.Disabled {
		return
	}
	switch event.Action {

	case Press:
		if handler.OnPress != nil {
			handler.OnPress(handler)
		}
		handler.pressed = true
		handler.selected = true
		focused = handler
		handler.Update()

	case Release:
		if handler.pressed {
			if handler.OnRelease != nil {
				handler.OnRelease(handler)
			}
			handler.pressed = false
			handler.Update()
		}
		fallthrough

	case Move:
		if !handler.hovered {
			if handler.OnHover != nil {
				handler.OnHover(handler)
			}
			hovered = append(hovered, handler)
			handler.hovered = true
			handler.Update()
		}
	}
}

func (handler *Handler) Rasterize() {
	handler.Style.Rasterize(handler.Size())
	handler.Update()
}

func (handler *Handler) Draw(target pixel.Target) {
	handler.background.Draw(target, handler.Center())
	handler.foreground.Draw(target, handler.Center())
}

func (handler *Handler) Dehover() {
	handler.pressed = false
	handler.hovered = false
	handler.Update()
}

func (handler *Handler) Defocus() {
	if handler == nil {
		return
	}
	handler.selected = false
	handler.Update()
}

func (handler *Handler) Update() {
	switch {
	case handler.pressed:
		handler.background = handler.Style.Active
	case handler.hovered:
		handler.background = handler.Style.Hover
	default:
		handler.background = handler.Style.Idle
	}
	switch {
	case handler.selected:
		handler.foreground = handler.Style.Focus
	case handler.Disabled:
		handler.foreground = handler.Style.Disabled
	default:
		handler.foreground = nil
	}
}

func (style HandlerStyle) Rasterize(w, h float64) {
	style.Idle.Rasterize(w, h)
	style.Hover.Rasterize(w, h)
	style.Active.Rasterize(w, h)
	style.Focus.Rasterize(w, h)
	style.Disabled.Rasterize(w, h)
}

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

func Focus(focus Focused) {
	if focused != nil {
		focused.Defocus()
	}
	focused = focus
}

var focused Focused
