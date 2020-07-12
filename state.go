package idle

import (
	"fmt"
	"image"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/fogleman/gg"
)

type (
	Styler interface {
		Rasterize(w, h float64)
		Draw(target pixel.Target, center pixel.Vec)
	}

	State struct {
		Fill   color.Color
		Drawer Drawer
		Border Border
		sprite *pixel.Sprite
	}

	Drawer interface {
		Draw(ctx *gg.Context)
	}

	Border struct {
		Color     color.Color
		Width     float64
		Inset     color.Color
		Alignment alignment
	}
)

func (style *State) Rasterize(w, h float64) {
	if !(w > 0 && h > 0) {
		panic(fmt.Errorf("incorrect tile size: %gx%g", w, h))
	}

	rgba := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	ctx := gg.NewContextForRGBA(rgba)

	if style.Fill != nil {
		ctx.SetColor(style.Fill)
		ctx.Clear()
	}

	if style.Drawer != nil {
		style.Drawer.Draw(ctx)
	}

	border := style.Border
	if border.Color != nil {
		if border.Width == 0 {
			border.Width = 1
		}
		ctx.SetLineWidth(border.Width * 2)
		if border.Alignment == 0 {
			border.Alignment = Left | Bottom | Right | Top
		}
		ctx.SetColor(border.Color)
		drawBorder(ctx, border.Alignment, w, h, 0)
		if border.Inset != nil {
			ctx.SetColor(border.Inset)
			ctx.SetLineWidth(1)
			drawBorder(ctx, border.Alignment, w, h, border.Width+0.5)
		}
	}

	if style.sprite == nil {
		style.sprite = pixel.NewSprite(pixel.PictureDataFromImage(rgba), pixel.R(0, 0, w, h))
	} else {
		style.sprite.Set(pixel.PictureDataFromImage(rgba), pixel.R(0, 0, w, h))
	}
}

func drawBorder(ctx *gg.Context, align alignment, w, h, p float64) {
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

func (style *State) Draw(target pixel.Target, center pixel.Vec) {
	style.sprite.Draw(target, pixel.IM.Moved(center))
}

type DrawerFunc func(ctx *gg.Context)

func (drawer DrawerFunc) Draw(ctx *gg.Context) {
	drawer(ctx)
}

type MultiDrawer []Drawer

func (drawer MultiDrawer) Draw(ctx *gg.Context) {
	for _, d := range drawer {
		d.Draw(ctx)
	}
}

type TextIconDrawer struct {
	TextColor color.Color
	IconColor color.Color
	Text      *Text
	Icon      *Icon
}

func (drawer *TextIconDrawer) Draw(ctx *gg.Context) {
	if drawer.Text != nil {
		ctx.SetColor(drawer.TextColor)
		drawer.Text.Draw(ctx)
	}
	if drawer.Icon != nil {
		ctx.SetColor(drawer.IconColor)
		drawer.Icon.Draw(ctx)
	}
}
