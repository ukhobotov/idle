package idle

import (
	"math"

	"github.com/faiface/pixel"
)

type (
	Location struct {
		Left, Bottom, Right, Top Constraint
		x1, y1, x2, y2           float64
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

func (loc *Location) Center() pixel.Vec {
	return pixel.V((loc.x1+loc.x2)/2, (loc.y1+loc.y2)/2)
}

func (loc *Location) Absolute() (x1, y1, x2, y2 float64) {
	return loc.x1, loc.y1, loc.x2, loc.y2
}

func (loc *Location) FitInto(x1, y1, x2, y2 float64) {
	loc.x1 = loc.Left.Absolute(x1, x2, false)
	loc.y1 = loc.Bottom.Absolute(y1, y2, false)
	loc.x2 = loc.Right.Absolute(x1, x2, true)
	loc.y2 = loc.Top.Absolute(y1, y2, true)
}

func (cst Constraint) Absolute(min, max float64, flipZero bool) float64 {
	if cst.Grid == 0 {
		if !flipZero {
			return min + math.Mod(max-min+cst.Margin, max-min)
		} else {
			return max - math.Mod(max-min-cst.Margin, max-min)
		}
	}
	return min + cst.Grid*(max-min) + cst.Margin
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
