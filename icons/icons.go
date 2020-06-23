package icons

import (
	"github.com/fogleman/gg"

	"carbon"
)

type Icon func(x, y, a float64) carbon.Drawing

const (
	l1 float64 = 3 / 32
	l2 float64 = 5 / 32
	l3 float64 = 7 / 32
	g3 float64 = 25 / 32
	g2 float64 = 27 / 32
	g1 float64 = 29 / 32
)

// var Back = func(style Style) carbon.Drawing {
//   iconColor, P, I := style.Props()
//   return func(ctx *gg.Context) {
//     ctx.SetColor(iconColor)
//     ctx.SetLineWidth(1)
//     ctx.DrawLine(P+I*3/16+0.5, P+I/16+0.5, P+I*3/16+0.5, P+I*15/16-0.5)
//     ctx.DrawLine(P+I/2+0.5, P+I/2+0.5, P+I*7/8-0.5, P+I*7/8-0.5)
//     ctx.DrawLine(P+I/2+0.5, P+I/2-0.5, P+I*7/8-0.5, P+I*1/8+0.5)
//     ctx.Stroke()
//   }
// }
//
// var Menu = func(style Style) carbon.Drawing {
//   iconColor, P, I := style.Props()
//   return func(ctx *gg.Context) {
//     ctx.SetColor(iconColor)
//     ctx.SetLineWidth(1)
//     ctx.DrawLine(P+I*1/8+0.5, P+I*3/16, P+I*7/8-0.5, P+I*3/16)
//     ctx.DrawLine(P+I*1/8+0.5, P+I*3/8+0.5, P+I*7/8-0.5, P+I*3/8+0.5)
//     ctx.DrawLine(P+I*1/8+0.5, P+I*5/8-0.5, P+I*7/8-0.5, P+I*5/8-0.5)
//     ctx.DrawLine(P+I*1/8+0.5, P+I*13/16, P+I*7/8-0.5, P+I*13/16)
//     ctx.Stroke()
//   }
// }

//
// func Refresh(style Style) carbon.Drawing {
//   iconColor, P, I := style.Props()
//   return func(ctx *gg.Context) {
//     ctx.SetColor(iconColor)
//     ctx.SetLineWidth(1)
//     ctx.DrawLine(P+I/8+0.5, P+I/8+0.5, P+I/8+0.5, P+I/4+0.5)
//     ctx.DrawLine(P+I/8+0.5, P+I/4+0.5, P+I/4+0.5, P+I/4+0.5)
//     ctx.DrawLine(P+I*3/4-0.5, P+I*3/4-0.5, P+I*7/8-0.5, P+I*3/4-0.5)
//     ctx.DrawLine(P+I*7/8-0.5, P+I*3/4-0.5, P+I*7/8-0.5, P+I*7/8-0.5)
//     ctx.Stroke()
//     ctx.DrawArc(P+I/2, P+I/2, I*3/8+0.5, 0, -math.Pi+0.56)
//     ctx.Stroke()
//     ctx.DrawArc(P+I/2, P+I/2, I*3/8+0.5, 0.56, math.Pi)
//     ctx.Stroke()
//   }
// }

func Play(x, y, a float64) carbon.Drawing {
	return func(ctx *gg.Context) {
		ctx.NewSubPath()
		ctx.MoveTo(x+a*l3, y+a*l1)
		ctx.LineTo(x+a*l3, y+a*g1)
		ctx.LineTo(x+a*29/32, y+a/2)
		ctx.ClosePath()
		ctx.SetLineWidth(a / 16)
		ctx.Stroke()
	}
}

func Stop(x, y, a float64) carbon.Drawing {
	return func(ctx *gg.Context) {
		ctx.SetLineWidth(a / 16)
		ctx.DrawRectangle(x+a*l2, y+a*l2, a*11/16, a*11/16)
		ctx.Stroke()
	}
}

func Pause(x, y, a float64) carbon.Drawing {
	return func(ctx *gg.Context) {
		ctx.SetLineWidth(a / 16)
		ctx.DrawRectangle(x+a*l2, y+a*l2, a*3/16, a*11/16)
		ctx.DrawRectangle(x+a*21/32, y+a*l2, a*3/16, a*11/16)
		ctx.Stroke()
	}
}

func Down(x, y, a float64) carbon.Drawing {
	return func(ctx *gg.Context) {
		ctx.MoveTo(x+a*5/32, y+a*11/32)
		ctx.LineTo(x+a/2, y+a*11/16)
		ctx.LineTo(x+a*27/32, y+a*11/32)
		ctx.SetLineWidth(a / 16)
		ctx.Stroke()
	}
}

func Exit(x, y, a float64) carbon.Drawing {
	return func(ctx *gg.Context) {
		ctx.DrawLine(x+a*g3, y+a*l1, x+a*g3, y+a*g1)
		ctx.DrawLine(x+a*l1, y*a*1/2, x*a*(l1+1/8), y*a*(1/2+1/8))
		ctx.DrawLine(x+a*l1, y*a*1/2, x*a*(l1+1/8), y*a*(1/2-1/8))
		ctx.DrawLine(x+a*l1, y*a*1/2, x*a*(g3-1/8), y*a*1/2)
		ctx.SetLineWidth(a / 16)
		ctx.Stroke()
	}
}

func Add(x, y, a float64) carbon.Drawing {
	return func(ctx *gg.Context) {
		ctx.SetLineWidth(a / 8)
		ctx.DrawLine(x+a*3/16, y+a/2, x+a*13/16, y+a/2)
		ctx.DrawLine(x+a/2, y+a*3/16, x+a/2, y+a*13/16)
		ctx.Stroke()
	}
}
