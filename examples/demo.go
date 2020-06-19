package main

import (
	"carbon"
)

func main() {
	carbon.Width, carbon.Height = 600, 400
	carbon.Fullscreen = false
	carbon.Root = &carbon.Container{
		Content: []carbon.Element{
			&carbon.Label{
				Layout: carbon.Layout{
					Left:   carbon.Grid(0.25),
					Bottom: carbon.Grid(0.25),
				},
				Color:    carbon.Text2,
				Text:     "Hello, carbon!",
				TextSize: 16,
			},
		},
	}
	carbon.Run()
}
