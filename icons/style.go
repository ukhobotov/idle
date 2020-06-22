package icons

type Style struct {
	Size   float64
	X0, Y0 float64
}

func (style Style) Props() (x, y, size float64) {
	return style.X0, style.Y0, style.Size
}
