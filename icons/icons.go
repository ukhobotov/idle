package icons

import (
	"github.com/fogleman/gg"

	"carbon"
)

type Icon func(style Style) carbon.Drawing

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

func Play(style Style) carbon.Drawing {
	x, y, I := style.Props()
	return func(ctx *gg.Context) {
		ctx.NewSubPath()
		ctx.MoveTo(x+I*l3, y+I*l1)
		ctx.LineTo(x+I*l3, y+I*g1)
		ctx.LineTo(x+I*29/32, y+I/2)
		ctx.ClosePath()
		ctx.SetLineWidth(I / 16)
		ctx.Stroke()
	}
}

func Stop(style Style) carbon.Drawing {
	x, y, I := style.Props()
	return func(ctx *gg.Context) {
		ctx.SetLineWidth(I / 16)
		ctx.DrawRectangle(x+I*l2, y+I*l2, I*11/16, I*11/16)
		ctx.Stroke()
	}
}

func Pause(style Style) carbon.Drawing {
	x, y, I := style.Props()
	return func(ctx *gg.Context) {
		ctx.SetLineWidth(I / 16)
		ctx.DrawRectangle(x+I*l2, y+I*l2, I*3/16, I*11/16)
		ctx.DrawRectangle(x+I*21/32, y+I*l2, I*3/16, I*11/16)
		ctx.Stroke()
	}
}

func Down(style Style) carbon.Drawing {
	x, y, I := style.Props()
	return func(ctx *gg.Context) {
		ctx.MoveTo(x+I*5/32, y+I*11/32)
		ctx.LineTo(x+I/2, y+I*11/16)
		ctx.LineTo(x+I*27/32, y+I*11/32)
		ctx.SetLineWidth(I / 16)
		ctx.Stroke()
	}
}

func Exit(style Style) carbon.Drawing {
	x, y, I := style.Props()
	return func(ctx *gg.Context) {
		ctx.DrawLine(x+I*g3, y+I*l1, x+I*g3, y+I*g1)
		ctx.DrawLine(x+I*l1, y*I*1/2, x*I*(l1+1/8), y*I*(1/2+1/8))
		ctx.DrawLine(x+I*l1, y*I*1/2, x*I*(l1+1/8), y*I*(1/2-1/8))
		ctx.DrawLine(x+I*l1, y*I*1/2, x*I*(g3-1/8), y*I*1/2)
		ctx.SetLineWidth(I / 16)
		ctx.Stroke()
	}
}

func Add(style Style) carbon.Drawing {
	x, y, I := style.Props()
	return func(ctx *gg.Context) {
		ctx.SetLineWidth(I / 8)
		ctx.DrawLine(x+I*3/16, y+I/2, x+I*13/16, y+I/2)
		ctx.DrawLine(x+I/2, y+I*3/16, x+I/2, y+I*13/16)
		ctx.Stroke()
	}
}
