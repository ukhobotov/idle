package carbon

import (
	"fmt"
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
)

type Text struct {
	Location
	Color     *color.RGBA
	Text      string
	TextSize  float64
	Font      *truetype.Font
	Alignment alignment

	matrix pixel.Matrix
	text   *text.Text
}

func (txt *Text) Handle(Event, float64, float64) {}

func (txt *Text) Rasterize() {
	x, y := txt.x1, txt.y1
	point := pixel.V(x, y)
	switch txt.Alignment {
	case 0:
		fallthrough
	case Bottom:
		txt.matrix = pixel.IM
	case Left:
		txt.matrix = pixel.IM.Rotated(point, -math.Pi/2)
	case Right:
		txt.matrix = pixel.IM.Rotated(point, math.Pi/2)
	case Top:
		txt.matrix = pixel.IM.Rotated(point, math.Pi)
	default:
		panic(fmt.Sprintf("wrong text alignment: %v", txt.Alignment))
	}

	if txt.Font == nil {
		txt.Font = Regular
	}
	if txt.TextSize == 0 {
		txt.TextSize = 16
	}

	txt.text = text.New(point, text.NewAtlas(NewFace(txt.Font, txt.TextSize), text.ASCII))
	txt.Update()
}

func (txt *Text) Update() {
	if *txt.Color == Transparent {
		txt.Color = Text1
	}
	txt.text.Color = txt.Color
	_, _ = txt.text.WriteString(txt.Text)
}

func (txt *Text) Draw(win *Window) {
	txt.text.Draw(win.window, txt.matrix)
}
