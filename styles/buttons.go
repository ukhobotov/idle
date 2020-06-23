package styles

import (
	"image/color"

	"github.com/fogleman/gg"

	"carbon"
)

var (
	TertiaryButton = carbon.HandlerStyle{
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
		Focused: ButtonFocus,
	}

	ButtonFocus = &carbon.Style{
		Border: carbon.Border{
			Color: carbon.Interactive4,
			Width: 2,
		},
	}

	ButtonDisabled = &carbon.Style{
		Fill:    &color.RGBA{},
		Drawing: nil,
		Border:  carbon.Border{},
	}
)
