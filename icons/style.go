package icons

import (
	"image/color"
)

type Style struct {
	Color  color.RGBA
	Size   float64
	X0, Y0 float64
}

func (style Style) Props() (clr color.RGBA, x, y, size float64) {
	return style.Color, style.X0, style.Y0, style.Size
}
