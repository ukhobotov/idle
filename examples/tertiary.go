package main

import (
	"github.com/fogleman/gg"

	"github.com/usualhuman/carbon"
	"github.com/usualhuman/carbon/icons"
	"github.com/usualhuman/carbon/styles"
)

func main() {
	window := &carbon.Window{
		Width:  256,
		Height: 128,
		Root: &carbon.Container{
			Content: []carbon.Element{
				&carbon.Handler{
					Location: carbon.Location{
						Left:   carbon.Margin(16),
						Top:    carbon.Margin(-16),
						Right:  carbon.Margin(16 + 192),
						Bottom: carbon.Margin(-16 - 48),
					},
					Style: styles.TertiaryButton,
					Final: func(ctx *gg.Context) {
						ctx.SetFontFace(carbon.NewFace(carbon.Regular, 16))
						ctx.DrawString("Look I'm tertiary!", 16, 48-19)
						icons.Add(float64(ctx.Width())-32, 16, 16)(ctx)
					},
				},
			},
		},
	}
	window.Show()
}
