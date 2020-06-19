package carbon

type (
	// Location contains an object box's position data
	Location struct {
		Left, Bottom, Right, Top Constraint // relative
		x1, y1, x2, y2           float64    // absolute
	}

	Constraint struct {
		Grid   float64
		Margin float64
	}
)

func (loc *Location) Size() (w, h float64) {
	return loc.x2 - loc.x1, loc.y2 - loc.y1
}

func (loc *Location) Contains(x, y float64) bool {
	return loc.x1 <= x && x < loc.x2 && loc.y1 <= y && y < loc.y2
}

func (loc *Location) Center() (x, y float64) {
	return (loc.x1 + loc.x2) / 2, (loc.y1 + loc.y2) / 2
}

func (loc *Location) FitInto(x1, y1, x2, y2 float64) {
	loc.x1 = loc.Left.Absolute(x1, x2, x1)
	loc.y1 = loc.Bottom.Absolute(y1, y2, y1)
	loc.x2 = loc.Right.Absolute(x1, x2, x2)
	loc.y2 = loc.Top.Absolute(y1, y2, y2)
}

func (loc *Location) Absolute() (x1, y1, x2, y2 float64) {
	return loc.x1, loc.y1, loc.x2, loc.y2
}

func (cst Constraint) Absolute(min, max, dft float64) float64 {
	delta := cst.Grid*(max-min) + cst.Margin
	switch {
	case min+delta > 0:
		return min + delta
	case min+delta < 0:
		return max + delta
	default:
		return dft
	}
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
