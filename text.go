package idle

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

type Text struct {
	Pos  pixel.Vec
	Font *truetype.Font
	Size float64
	Text string
}

func (text *Text) Draw(ctx *gg.Context) {
	if text.Font == nil {
		text.Font = Regular
	}
	if text.Size == 0 {
		text.Size = 16
	}
	ctx.SetFontFace(NewFace(text.Font, text.Size))
	w, h := float64(ctx.Width()), float64(ctx.Height())
	ctx.DrawString(text.Text, math.Mod(text.Pos.X+w, w), h-math.Mod(text.Pos.Y+h, h))
}
