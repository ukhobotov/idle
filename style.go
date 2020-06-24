package carbon

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/fogleman/gg"
)

type Style struct {
	Fill    *color.RGBA
	Drawing Drawing
	Border  Border

	sprite *pixel.Sprite
}

type Drawing func(ctx *gg.Context)

type Border struct {
	Color     *color.RGBA
	Width     float64
	Alignment alignment
}

func (style *Style) Rasterize(w, h float64) {
	if style == nil {
		return
	}
	if !(w > 0 && h > 0) {
		panic(fmt.Errorf("incorrect size: %g, %g", w, h))
	}
	if style.Fill == nil && style.Drawing == nil && style.Border == (Border{}) {
		return
	}

	ctx := gg.NewContext(int(w), int(h))

	if style.Fill != nil {
		ctx.SetColor(style.Fill)
		ctx.DrawRectangle(0, 0, w, h)
		ctx.Fill()
	}

	if style.Drawing != nil {
		style.Drawing(ctx)
	}

	if style.Border.Color != nil {
		if style.Border.Width == 0 {
			style.Border.Width = 1
		}
		ctx.SetLineWidth(style.Border.Width * 2)
		if style.Border.Alignment == 0 {
			style.Border.Alignment = Left | Bottom | Right | Top
		}
		ctx.SetColor(style.Border.Color)
		if style.Border.Alignment.Has(Left) {
			ctx.DrawLine(0, h, 0, 0)
			ctx.Stroke()
		}
		if style.Border.Alignment.Has(Top) {
			ctx.DrawLine(0, 0, w, 0)
			ctx.Stroke()
		}
		if style.Border.Alignment.Has(Right) {
			ctx.DrawLine(w, 0, w, h)
			ctx.Stroke()
		}
		if style.Border.Alignment.Has(Bottom) {
			ctx.DrawLine(w, h, 0, h)
			ctx.Stroke()
		}
	}

	if style.sprite == nil {
		style.sprite = pixel.NewSprite(pixel.PictureDataFromImage(ctx.Image()), pixel.R(0, 0, w, h))
	} else {
		style.sprite.Set(pixel.PictureDataFromImage(ctx.Image()), pixel.R(0, 0, w, h))
	}
}

func (style *Style) Draw(win *Window, x, y float64) {
	if style == nil || style.sprite == nil {
		return
	}
	style.sprite.Draw(win.window, pixel.IM.Moved(pixel.V(x, y)))
}

func (style *Style) Empty() bool {
	return style.sprite == nil
}

type Styler interface {
	Rasterize(w, h float64)
	Draw(x, y float64)
}
