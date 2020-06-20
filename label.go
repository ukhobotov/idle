package carbon

import (
	"fmt"
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
)

const defaultTextSize = 16

type Text struct {
	Location
	Color     color.RGBA
	Text      string
	TextSize  float64
	Font      *truetype.Font
	Alignment alignment
	matrix    pixel.Matrix
	text      *text.Text
}

func (label *Text) Handle(Event, float64, float64) {}

func (label *Text) Rasterize() {
	x, y := label.x1, label.y1
	point := pixel.V(x, y)
	switch label.Alignment {
	case 0:
		fallthrough
	case Bottom:
		label.matrix = pixel.IM
	case Left:
		label.matrix = pixel.IM.Rotated(point, -math.Pi/2)
	case Right:
		label.matrix = pixel.IM.Rotated(point, math.Pi/2)
	case Top:
		label.matrix = pixel.IM.Rotated(point, math.Pi)
	default:
		panic(fmt.Sprintf("wrong text alignment: %v", label.Alignment))
	}

	if label.Font == nil {
		label.Font = Regular
	}
	if label.TextSize == 0 {
		label.TextSize = defaultTextSize
	}
	if label.Color == Transparent {
		label.Color = Text1
	}

	label.text = text.New(point, text.NewAtlas(NewFace(label.Font, label.TextSize), text.ASCII))
	label.text.Color = label.Color
	_, _ = label.text.WriteString(label.Text)
}

func (label *Text) Draw(win *Window) {
	label.text.Draw(win.window, label.matrix)
}
