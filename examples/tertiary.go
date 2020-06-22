package main

import (
	"github.com/fogleman/gg"

	"carbon"
	"carbon/icons"
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
					Style: carbon.HandlerStyle{
						Idle: &carbon.Style{
							Border: carbon.Border{Color: carbon.Interactive3},
							Drawing: func(ctx *gg.Context) {
								ctx.SetColor(carbon.Text4)
							},
						},
						Hover: &carbon.Style{
							Fill: carbon.HoverTertiary,
							Drawing: func(ctx *gg.Context) {
								ctx.SetColor(carbon.Inverse1)
							},
						},
						Active: &carbon.Style{
							Fill: carbon.ActiveTertiary,
							Drawing: func(ctx *gg.Context) {
								ctx.SetColor(carbon.Inverse1)
							},
						},
					},
					Common: func(ctx *gg.Context) {
						ctx.SetFontFace(carbon.NewFace(carbon.Regular, 16))
						ctx.DrawString("Look I'm tertiary!", 16, 48-19)
						icons.Add(icons.Style{
							Size: 16,
							X0:   192 - 16 - 16,
							Y0:   16,
						})(ctx)
					},
				},
			},
		},
	}
	window.Show()
}
