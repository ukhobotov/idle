package components

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/usualhuman/idle"
)

type (

	// Button is element used primarily for actions. There are some kinds of it that might be used for special purposes.
	Button struct {
		idle.Location
		Style    *ButtonStyle
		Text     *idle.Text
		Icon     *idle.Icon
		Click    func(this *Button, event idle.Event)
		Disabled bool

		pressed bool
		hovered bool
		focused bool
	}

	ButtonStyle struct {
		Idle     *idle.State
		Hover    *idle.State
		Active   *idle.State
		Focus    *idle.State
		Disabled *idle.State
	}
)

func (bs *ButtonStyle) States() []*idle.State {
	return []*idle.State{bs.Idle, bs.Hover, bs.Active, bs.Focus, bs.Disabled}
}

func (button *Button) Rasterize() {
	for _, state := range button.Style.States() {
		if state.Drawer != nil {
			switch drawer := state.Drawer.(type) {
			case *idle.TextIconDrawer:
				drawer.Text = button.Text
				drawer.Icon = button.Icon
			}
		}
		state.Rasterize(button.Size())
	}
}

func (button *Button) Handle(event idle.Event) {
	if button.Disabled {
		return
	}
	switch event.Action {
	case idle.Press:
		if event.Button == pixelgl.MouseButtonLeft {
			button.pressed = true
			button.focused = true
			idle.Focus(button)
		}
	case idle.Release:
		if event.Button == pixelgl.MouseButtonLeft {
			if button.Click != nil {
				button.Click(button, event)
			}
			button.pressed = false
		}
		fallthrough
	case idle.Move:
		if !button.hovered {
			button.hovered = true
			idle.Hover(button)
		}
	}
}

func (button *Button) Draw(target pixel.Target) {
	if button.Disabled {
		button.Style.Disabled.Draw(target, button.Center())
		return
	}
	switch {
	case button.pressed:
		button.Style.Active.Draw(target, button.Center())
	case button.hovered:
		button.Style.Hover.Draw(target, button.Center())
	default:
		button.Style.Idle.Draw(target, button.Center())
	}
	if button.focused {
		button.Style.Focus.Draw(target, button.Center())
	}
}

func (button *Button) Dehover() {
	button.hovered = false
}

func (button *Button) Defocus() {
	button.focused = false
}
