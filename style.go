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
	Splitter  *color.RGBA
	Alignment alignment
}

func (style *Style) Rasterize(w, h float64) {
	if style == nil {
		return
	}
	if !(w > 0 && h > 0) {
		panic(fmt.Errorf("incorrect style size: %g, %g", w, h))
	}

	ctx := gg.NewContext(int(w), int(h))

	if style.Fill != nil {
		ctx.SetColor(style.Fill)
		ctx.Clear()
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
		drawBorder(style.Border.Alignment, ctx, w, h, 0)
		if style.Border.Splitter != nil {
			ctx.SetColor(style.Border.Splitter)
			ctx.SetLineWidth(1)
			drawBorder(style.Border.Alignment, ctx, w, h, style.Border.Width+0.5)
		}
	}

	if style.sprite == nil {
		style.sprite = pixel.NewSprite(pixel.PictureDataFromImage(ctx.Image()), pixel.R(0, 0, w, h))
	} else {
		style.sprite.Set(pixel.PictureDataFromImage(ctx.Image()), pixel.R(0, 0, w, h))
	}
}

func drawBorder(align alignment, ctx *gg.Context, w, h, p float64) {
	if align.Has(Left) {
		ctx.DrawLine(0+p, h-p, 0+p, 0+p)
		ctx.Stroke()
	}
	if align.Has(Top) {
		ctx.DrawLine(0+p, 0+p, w-p, 0+p)
		ctx.Stroke()
	}
	if align.Has(Right) {
		ctx.DrawLine(w-p, 0+p, w-p, h-p)
		ctx.Stroke()
	}
	if align.Has(Bottom) {
		ctx.DrawLine(w-p, h-p, 0+p, h-p)
		ctx.Stroke()
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
