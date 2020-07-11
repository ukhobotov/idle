package idle

import (
	"fmt"
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
)

type Label struct {
	Location
	Color     color.Color
	Text      string
	TextSize  float64
	Font      *truetype.Font
	Alignment alignment

	matrix pixel.Matrix
	text   *text.Text
}

func (label *Label) Handle(Event) {}

func (label *Label) Rasterize() {
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
		label.TextSize = 16
	}
	if label.text == nil {
		label.text = text.New(point, text.NewAtlas(NewFace(label.Font, label.TextSize), text.ASCII))
	}

	label.text.Color = label.Color
	_, _ = label.text.WriteString(label.Text)
}

func (label *Label) Draw(target pixel.Target) {
	label.text.Draw(target, label.matrix)
}
