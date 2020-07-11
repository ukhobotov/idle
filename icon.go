package idle

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/fogleman/gg"
)

type (
	Icon struct {
		Pos   pixel.Vec
		Side  float64
		Color color.Color
		Func  IconFunc
	}

	IconFunc func(x, y, a float64) DrawerFunc
)

func (icon *Icon) Draw(ctx *gg.Context) {
	if icon.Color != nil {
		ctx.SetColor(icon.Color)
	}
	w, h := float64(ctx.Width()), float64(ctx.Height())
	icon.Func(math.Mod(icon.Pos.X+w, w), h-icon.Side-math.Mod(icon.Pos.Y+h, h), icon.Side)(ctx)
}
