package components

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/usualhuman/idle"
	"github.com/usualhuman/idle/carbon"
)

type TextInput struct {
	idle.Location
	Style       *InputStyle
	Placeholder string
	Entered     *string
	Disabled    bool
	focused     bool
}

func (input *TextInput) Rasterize() {
	for _, state := range input.Style.States() {
		state.Rasterize(input.Size())
	}
}

func (input *TextInput) Handle(event idle.Event) {
	switch event.Action {
	case idle.Press:
		if event.Button == pixelgl.MouseButtonLeft {
			input.focused = true
			idle.Focus(input)
		}
	}
}

func (input *TextInput) Draw(target pixel.Target) {
	input.Style.Draw(target, input.Center(), input.focused, input.Disabled)
}

func (input *TextInput) Defocus() {
	input.focused = false
}

type InputStyle struct {
	Idle     *idle.State
	Focus    *idle.State
	Error    *idle.State
	Disabled *idle.State
}

func (style *InputStyle) States() []*idle.State {
	return []*idle.State{style.Idle, style.Focus, style.Disabled}
}

func (style *InputStyle) Rasterize(w, h float64) {
	for _, state := range []*idle.State{style.Idle, style.Focus, style.Disabled} {
		state.Rasterize(w, h)
	}
}

func (style *InputStyle) Draw(target pixel.Target, center pixel.Vec, focused, disabled bool) {
	if disabled {
		style.Disabled.Draw(target, center)
		return
	}
	style.Idle.Draw(target, center)
	if focused {
		style.Focus.Draw(target, center)
	}
}

func Input() *InputStyle {
	return &InputStyle{
		Idle: &idle.State{
			Fill:   carbon.Field1,
			Drawer: &idle.TextIconDrawer{},
			Border: idle.Border{
				Color:     carbon.Ui4,
				Alignment: idle.Bottom,
			},
		},
		Focus: &idle.State{
			Border: idle.Border{
				Color: carbon.Focus,
				Width: 2,
			},
		},
		Disabled: &idle.State{
			Fill:   carbon.Disabled2,
			Drawer: &idle.TextIconDrawer{},
		},
	}
}
