package idle

import (
	"image/color"
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
	*pixelgl.Window
}

func (win *Window) Show() {
	shown = append(shown, win)
	pixelgl.Run(func() {
		var (
			err  error
			pos0 pixel.Vec
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
		win.Window, err = pixelgl.NewWindow(pixelgl.WindowConfig{
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
		for !win.Window.Closed() {
			win.Window.Update()

			// handling resizing
			if w1, h1 := win.Window.Bounds().Size().XY(); w1 != win.Width || h1 != win.Height {
				win.Width, win.Height = w1, h1
				// TODO: add minsize (why?)
				win.Update()
			}

			// transmitting mouse events
			pos := win.Window.MousePosition()
			buttons := []pixelgl.Button{
				pixelgl.MouseButtonLeft,
				pixelgl.MouseButtonRight,
				pixelgl.MouseButtonMiddle,
			}
			pressed := NilButton
			for _, button := range buttons {
				if win.Window.Pressed(button) {
					pressed = button
				}
				if win.Window.JustPressed(button) {
					if focused != nil {
						focused.Defocus()
						focused = nil
					}
					win.Root.Handle(Press.The(button, pos))
				}
				if win.Window.JustReleased(button) {
					win.Root.Handle(Release.The(button, pos))
					HandleHovered(Release.The(pressed, pos))
				}
			}
			if pos != pos0 {
				win.Root.Handle(Move.The(pressed, pos))
				HandleHovered(Move.The(pressed, pos))
				pos0 = pos
			}
			scroll := win.Window.MouseScroll()
			if scroll != pixel.ZV {
				win.Root.Handle(Event{Action: Scroll, Scroll: scroll, MousePos: pos})
			}

			// counting fps
			end := time.Now()
			frames++
			if end.Sub(start).Seconds() > 1 {
				start = end
				frames = 0
			}

			// drawing the UI
			win.Window.Clear(color.RGBA{})
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
