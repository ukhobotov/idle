package carbon

import (
	"github.com/faiface/pixel/pixelgl"
)

type (
	Handler struct {
		Location
		Background     HandlerStyle
		Foreground     *Style
		Press, Release func(this *Handler)

		background       *Style
		hovered, pressed bool
	}

	HandlerStyle struct {
		Idle, Hover, Active *Style
	}

	Hover interface {
		Contains(x, y float64) bool
		Land()
	}
)

func (handler *Handler) Handle(event Event, _, _ float64) {
	switch event {

	case Press.The(pixelgl.MouseButtonLeft):
		if handler.Press != nil {
			handler.Press(handler)
		}

		handler.pressed = true
		handler.update()

	case Release.The(pixelgl.MouseButtonLeft):
		if handler.Release != nil {
			handler.Release(handler)
		}
		handler.hovered = true
		handler.pressed = false
		handler.update()

	case MoveEvent:
		if !handler.hovered {
			hovered = append(hovered, handler)
			handler.hovered = true
			handler.update()
		}
	}
}

func (handler *Handler) Rasterize() {
	handler.Background.Rasterize(handler.Size())
	handler.Foreground.Rasterize(handler.Size())
	handler.update()
}

func (handler *Handler) Draw() {
	handler.background.Draw(handler.Center())
	handler.Foreground.Draw(handler.Center())
}

func (handler *Handler) Land() {
	handler.pressed = false
	handler.hovered = false
	handler.update()
}

func (handler *Handler) update() {
	switch {
	case handler.pressed:
		handler.background = handler.Background.Active
	case handler.hovered:
		handler.background = handler.Background.Hover
	default:
		handler.background = handler.Background.Idle
	}
}

func (style *HandlerStyle) Rasterize(w, h float64) {
	style.Idle.Rasterize(w, h)
	style.Hover.Rasterize(w, h)
	style.Active.Rasterize(w, h)
}

var hovered []*Handler

func HandleHovered(event Event, x, y float64) {
	for _, hover := range hovered {
		if !hover.Contains(x, y) {
			hover.Land()
		}
	}
}

var (
	ButtonPrimary = HandlerStyle{
		Idle: &Style{
			Fill: Interactive1,
		},
		Hover: &Style{
			Fill: HoverPrimary,
		},
		Active: &Style{
			Fill: ActivePrimary,
		},
	}
	ButtonPrimaryText = HandlerStyle{
		Idle: &Style{
			Fill: Interactive1,
		},
		Hover: &Style{
			Fill: HoverPrimaryText,
		},
		Active: &Style{
			Fill: ActivePrimary,
		},
	}
	ButtonSecondary = HandlerStyle{
		Idle: &Style{
			Fill: Interactive2,
		},
		Hover: &Style{
			Fill: HoverSecondary,
		},
		Active: &Style{
			Fill: ActiveSecondary,
		},
	}
	ButtonTertiary = HandlerStyle{
		Idle: &Style{
			Border: Border{
				Color: Interactive3,
			},
		},
		Hover: &Style{
			Fill: HoverTertiary,
		},
		Active: &Style{
			Fill: ActiveTertiary,
		},
	}
)
