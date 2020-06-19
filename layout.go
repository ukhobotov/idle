package carbon

type (
	// Layout contains an object box's position data
	Layout struct {
		Left, Bottom, Right, Top Constraint // relative
		x1, y1, x2, y2           float64    // absolute
	}

	// Constraint is a way of bound definition. Can be a Margin, a Grid
	Constraint interface {
		Absolute(min, max float64) float64
	}
	Margin     float64
	Grid       float64
	Absolute   float64
	gridOffset struct {
		Grid   float64
		Offset float64
	}
)

func (layout *Layout) Size() (w, h float64) {
	return layout.x2 - layout.x1, layout.y2 - layout.y1
}

func (layout *Layout) Contains(x, y float64) bool {
	return layout.x1 <= x && x < layout.x2 && layout.y1 <= y && y < layout.y2
}

func (layout *Layout) Center() (x, y float64) {
	return (layout.x1 + layout.x2) / 2, (layout.y1 + layout.y2) / 2
}

func (layout *Layout) FitInto(x1, y1, x2, y2 float64) {
	layout.x1 = absolute(layout.Left, x1, x2, x1)
	layout.y1 = absolute(layout.Bottom, y1, y2, y1)
	layout.x2 = absolute(layout.Right, x1, x2, x2)
	layout.y2 = absolute(layout.Top, y1, y2, y2)
}

func absolute(cst Constraint, min, max, dft float64) float64 {
	if cst != nil {
		return cst.Absolute(min, max)
	} else {
		return dft
	}
}

func (layout *Layout) Absolute() (x1, y1, x2, y2 float64) {
	return layout.x1, layout.y1, layout.x2, layout.y2
}

func (margin Margin) Absolute(min, max float64) float64 {
	if margin >= 0 {
		return min + float64(margin)
	} else {
		return max + float64(margin)
	}
}

func (grid Grid) Absolute(min, max float64) float64 {
	return min + float64(grid)*(max-min)
}

func (abs Absolute) Absolute(float64, float64) float64 {
	return float64(abs)
}

func GridOffset(grid, offset float64) Constraint {
	return gridOffset{
		Grid:   grid,
		Offset: offset,
	}
}

func (gdOs gridOffset) Absolute(min, max float64) float64 {
	return min + gdOs.Grid*(max-min) + gdOs.Offset
}
