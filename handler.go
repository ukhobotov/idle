package carbon

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/fogleman/gg"
)

type (
	Handler struct {
		Location
		Style  HandlerStyle
		Common Drawing

		OnPress, OnRelease, OnHover func(this *Handler)

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
	switch {

	case event == Press.The(pixelgl.MouseButtonLeft):
		if handler.OnPress != nil {
			handler.OnPress(handler)
		}
		handler.pressed = true
		handler.Update()

	case event == Release.The(pixelgl.MouseButtonLeft):
		if handler.pressed {
			if handler.OnRelease != nil {
				handler.OnRelease(handler)
			}
			handler.pressed = false
			handler.Update()
		}
		fallthrough

	case event.Action == Move:
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
	handler.Style.Prepare(handler.Common)
	handler.Style.Rasterize(handler.Size())
	handler.Update()
}

func (handler *Handler) Draw(win *Window) {
	x, y := handler.Center()
	handler.background.Draw(win, x, y)
}

func (handler *Handler) Land() {
	handler.pressed = false
	handler.hovered = false
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
}

func (style *HandlerStyle) Prepare(common Drawing) {
	for _, style := range []*Style{style.Idle, style.Hover, style.Active} {
		local := style.Drawing
		style.Drawing = func(ctx *gg.Context) {
			local(ctx)
			common(ctx)
		}
	}
}

func (style *HandlerStyle) Rasterize(w, h float64) {
	style.Idle.Rasterize(w, h)
	style.Hover.Rasterize(w, h)
	style.Active.Rasterize(w, h)
}

var hovered []*Handler

func HandleHovered(event Event, x, y float64) {
	if event.Action == Move && event.Button != NilButton {
		return
	}
	for i, hover := range hovered {
		if !hover.Contains(x, y) {
			hover.Land()
			last := len(hovered) - 1
			hovered[i], hovered[last] = hovered[last], nil
			hovered = hovered[:last]
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
