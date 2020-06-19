package main

import (
	"carbon"
)

func main() {
	carbon.Width, carbon.Height = 256, 128
	carbon.Fullscreen = false
	carbon.Root = &carbon.Container{
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
	}
	carbon.Run()
}
