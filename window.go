package carbon

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	Title         string
	Width, Height float64
	Fullscreen    bool
	Root          Element

	window *pixelgl.Window
}

func (win *Window) Show() {
	shown = append(shown, win)
	pixelgl.Run(func() {
		var (
			err    error
			x0, y0 float64
		)

		if win.Fullscreen {
			glfw.WindowHint(glfw.Maximized, glfw.True)
		}
		if win.Width == 0 {
			win.Width = 1280
		}
		if win.Height == 0 {
			win.Height = 720
		}

		// creating a window
		win.window, err = pixelgl.NewWindow(pixelgl.WindowConfig{
			Title:     win.Title,
			Bounds:    pixel.R(0, 0, win.Width, win.Height),
			VSync:     true,
			Resizable: true,
		})
		if err != nil {
			panic(err)
		}

		win.Update()

		start := time.Now()
		frames := 0

		// entering the loop
		for !win.window.Closed() {
			win.window.Update()

			// handling resizing
			if w1, h1 := win.window.Bounds().Size().XY(); w1 != win.Width || h1 != win.Height {
				win.Width, win.Height = w1, h1
				// TODO: add minsize
				win.Update()
			}

			// transmitting mouse events
			x, y := win.window.MousePosition().XY()
			buttons := []pixelgl.Button{
				pixelgl.MouseButtonLeft,
				pixelgl.MouseButtonRight,
				pixelgl.MouseButtonMiddle,
			}
			pressed := NilButton
			for _, button := range buttons {
				if win.window.Pressed(button) {
					pressed = button
				}
				if win.window.JustPressed(button) {
					if focused != nil {
						focused.Defocus()
						focused = nil
					}
					win.Root.Handle(Press.The(button), x, y)
				}
				if win.window.JustReleased(button) {
					win.Root.Handle(Release.The(button), x, y)
					HandleHovered(Release.The(pressed), x, y)
				}
			}
			if x != x0 || y != y0 {
				win.Root.Handle(Move.The(pressed), x, y)
				HandleHovered(Move.The(pressed), x, y)
				x0, y0 = x, y
			}
			scroll := win.window.MouseScroll()
			if scroll != pixel.ZV {
				win.Root.Handle(Event{Action: Scroll, Scroll: scroll}, x, y)
			}

			// counting fps
			end := time.Now()
			frames++
			if end.Sub(start).Seconds() > 1 {
				start = end
				frames = 0
			}

			// drawing the UI
			win.window.Clear(Bg)
			win.Root.Draw(win)
		}
		for i := range shown {
			if shown[i] == win {
				last := len(shown) - 1
				shown[i], shown[last] = shown[last], nil
				shown = shown[:last]
			}
		}
	})
}

func (win *Window) Update() {
	win.Root.FitInto(0, 0, win.Width, win.Height)
	win.Root.Rasterize()
}

var shown []*Window

func UpdateShown() {
	for _, window := range shown {
		window.Update()
	}
}
