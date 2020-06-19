package carbon

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	window     *pixelgl.Window
	Height     float64 = 720
	Width      float64 = 1280
	Fullscreen         = true
)

var Root Element

func Run() {
	pixelgl.Run(func() {
		var (
			err    error
			x0, y0 float64
		)

		if Fullscreen {
			glfw.WindowHint(glfw.Maximized, glfw.True)
		}

		// create a window
		window, err = pixelgl.NewWindow(pixelgl.WindowConfig{
			Title:     "Four Suns",
			Bounds:    pixel.R(0, 0, Width, Height),
			VSync:     true,
			Resizable: true,
		})
		if err != nil {
			panic(err)
		}

		UpdateRoot()

		start := time.Now()
		frames := 0

		// enter the loop
		for !window.Closed() {
			window.Update()

			// handle resizing
			if w1, h1 := window.Bounds().Size().XY(); w1 != Width || h1 != Height {
				Width, Height = w1, h1
				UpdateRoot()
			}

			// transmit mouse events
			x, y := window.MousePosition().XY()
			buttons := []pixelgl.Button{
				pixelgl.MouseButtonLeft,
				pixelgl.MouseButtonRight,
				pixelgl.MouseButtonMiddle,
			}
			pressed := NilButton
			for _, button := range buttons {
				if window.Pressed(button) {
					pressed = button
				}
				if window.JustPressed(button) {
					Root.Handle(Press.The(button), x, y)
				}
				if window.JustReleased(button) {
					Root.Handle(Release.The(button), x, y)
				}
			}
			if x != x0 || y != y0 {
				Root.Handle(Move.The(pressed), x, y)
				HandleHovered(Move.The(pressed), x, y)
				x0, y0 = x, y
			}
			scroll := window.MouseScroll()
			if scroll != pixel.ZV {
				Root.Handle(Event{Action: Scroll, Scroll: scroll}, x, y)
			}

			// count fps
			end := time.Now()
			frames++
			if end.Sub(start).Seconds() > 1 {
				start = end
				frames = 0
			}

			// draw the UI
			window.Clear(Bg)
			Root.Draw()
		}
	})
}

func UpdateRoot() {
	Root.FitInto(0, 0, Width, Height)
	Root.Rasterize()
}
