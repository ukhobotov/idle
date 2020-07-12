package components

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/usualhuman/idle"
	"github.com/usualhuman/idle/carbon"
)

// Button is an element used primarily for actions. There are some kinds of it that might be used for special purposes.
type Button struct {
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

// ButtonStyle declares the button's look
type ButtonStyle struct {
	Idle     *idle.State
	Hover    *idle.State
	Active   *idle.State
	Focus    *idle.State
	Disabled *idle.State
}

func (bs *ButtonStyle) States() []*idle.State {
	return []*idle.State{bs.Idle, bs.Hover, bs.Active, bs.Focus, bs.Disabled}
}

// PrimaryButton is needed for the principal call to action on the page.
func PrimaryButton() *ButtonStyle {
	return &ButtonStyle{
		Idle: &idle.State{
			Fill: carbon.Interactive1,
			Drawer: &idle.TextIconDrawer{
				TextColor: carbon.Text4,
				IconColor: carbon.Icon3,
			},
		},
		Hover: &idle.State{
			Fill: carbon.HoverPrimary,
			Drawer: &idle.TextIconDrawer{
				TextColor: carbon.Text4,
				IconColor: carbon.Icon3,
			},
		},
		Active: &idle.State{
			Fill: carbon.ActivePrimary,
			Drawer: &idle.TextIconDrawer{
				TextColor: carbon.Text4,
				IconColor: carbon.Icon3,
			},
		},
		Focus: &idle.State{
			Border: idle.Border{
				Color: carbon.Focus,
				Width: 2,
				Inset: carbon.Inverse1,
			},
		},
		Disabled: &idle.State{
			Fill: carbon.Disabled2,
			Drawer: &idle.TextIconDrawer{
				TextColor: carbon.Disabled3,
				IconColor: carbon.Disabled3,
			},
		},
	}
}

// SecondaryButton is needed for secondary actions on each page.
func SecondaryButton() *ButtonStyle {
	return &ButtonStyle{
		Idle: &idle.State{
			Fill: carbon.Interactive2,
			Drawer: &idle.TextIconDrawer{
				TextColor: carbon.Text4,
				IconColor: carbon.Icon3,
			},
		},
		Hover: &idle.State{
			Fill: carbon.HoverSecondary,
			Drawer: &idle.TextIconDrawer{
				TextColor: carbon.Text4,
				IconColor: carbon.Icon3,
			},
		},
		Active: &idle.State{
			Fill: carbon.ActiveSecondary,
			Drawer: &idle.TextIconDrawer{
				TextColor: carbon.Text4,
				IconColor: carbon.Icon3,
			},
		},
		Focus: &idle.State{
			Border: idle.Border{
				Color: carbon.Focus,
				Width: 2,
				Inset: carbon.Inverse1,
			},
		},
		Disabled: &idle.State{
			Fill: carbon.Disabled2,
			Drawer: &idle.TextIconDrawer{
				TextColor: carbon.Disabled3,
				IconColor: carbon.Disabled3,
			},
		},
	}
}

// TertiaryButton is needed for some additional actions.
func TertiaryButton() *ButtonStyle {
	return &ButtonStyle{
		Idle: &idle.State{
			Border: idle.Border{
				Color: carbon.Interactive3,
			},
		},
		Hover: &idle.State{
			Fill: carbon.HoverTertiary,
		},
		Active: &idle.State{
			Fill: carbon.ActiveTertiary,
		},
	}
}
