package main

import (
	"github.com/faiface/pixel"
	"github.com/usualhuman/idle/carbon"
	"github.com/usualhuman/idle/carbon/components"
	"github.com/usualhuman/idle/carbon/icons"

	"github.com/usualhuman/idle"
)

func main() {
	window := &idle.Window{
		Width:  256,
		Height: 128,
		Root: &idle.Container{
			State: &idle.State{
				Fill: carbon.Background,
			},
			Content: []idle.Element{
				&components.Button{
					Location: idle.Location{
						Left:   idle.Margin(16),
						Top:    idle.Margin(-16),
						Right:  idle.Margin(-16),
						Bottom: idle.Margin(-16 - 48),
					},
					Style: carbon.PrimaryButton(),
					Text: &idle.Text{
						Pos:  pixel.V(16, 18),
						Text: "Look I'm renovated!",
					},
					Icon: &idle.Icon{
						Pos:  pixel.V(-32, 16),
						Side: 16,
						Func: icons.Add,
					},
				},
			},
		},
	}
	window.Show()
}
