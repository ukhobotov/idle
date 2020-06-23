package carbon

type (
	Location interface {
		Size() (w, h float64)
		Contains(x, y float64) bool
		Center() (x, y float64)
		FitInto(x1, y1, x2, y2 float64)
	}

	Bounds struct {
		Left, Bottom, Right, Top Constraint
		absolute
	}

	Fixed struct {
		Horiz, Vert   Constraint
		Width, Height float64
		absolute
	}

	Constraint struct {
		Grid   float64
		Margin float64
	}

	absolute struct {
		x1, y1, x2, y2 float64
	}
)

func (abs *absolute) Size() (w, h float64) {
	return abs.x2 - abs.x1, abs.y2 - abs.y1
}

func (abs *absolute) Contains(x, y float64) bool {
	return abs.x1 <= x && x < abs.x2 && abs.y1 <= y && y < abs.y2
}

func (abs *absolute) Center() (x, y float64) {
	return (abs.x1 + abs.x2) / 2, (abs.y1 + abs.y2) / 2
}

func (abs *absolute) Absolute() (x1, y1, x2, y2 float64) {
	return abs.x1, abs.y1, abs.x2, abs.y2
}

func (bounds *Bounds) FitInto(x1, y1, x2, y2 float64) {
	bounds.x1 = bounds.Left.Absolute(x1, x2, x1)
	bounds.y1 = bounds.Bottom.Absolute(y1, y2, y1)
	bounds.x2 = bounds.Right.Absolute(x1, x2, x2)
	bounds.y2 = bounds.Top.Absolute(y1, y2, y2)
}

func (cst Constraint) Absolute(min, max, dft float64) float64 {
	abs := min + cst.Grid*(max-min) + cst.Margin
	if abs == 0 {
		return dft
	}
	if abs < 0 {
		abs += max - min
	}
	return abs
}

func Margin(margin float64) Constraint {
	return Constraint{
		Margin: margin,
	}
}

func Grid(grid float64) Constraint {
	return Constraint{
		Grid: grid,
	}
}

func GridMargin(grid, margin float64) Constraint {
	return Constraint{
		Grid:   grid,
		Margin: margin,
	}
}
