package main

import (
	"carbon"
)

func main() {
	window := &carbon.Window{
		Height:     256,
		Width:      128,
		Fullscreen: true,
		Root: &carbon.Container{
			Content: []carbon.Element{
				&carbon.Text{
					Location: carbon.Location{
						Left:   carbon.Margin(32),
						Bottom: carbon.Margin(-32),
					},
					Color:    carbon.Text2,
					Text:     "Hello, carbon!",
					TextSize: 16,
				},
			},
		},
	}

	window.Show()
}
