package styles

import (
	"image/color"

	"github.com/fogleman/gg"

	"github.com/usualhuman/carbon"
)

var (
	PrimaryButton = carbon.HandlerStyle{
		Idle: &carbon.Style{
			Fill: carbon.Interactive1,
		},
		Hover: &carbon.Style{
			Fill: carbon.HoverPrimary,
		},
		Active: &carbon.Style{
			Fill: carbon.ActivePrimary,
		},
		Selected: ButtonSelected,
	}
	SecondaryButton = carbon.HandlerStyle{
		Idle: &carbon.Style{
			Fill: carbon.Interactive2,
		},
		Hover: &carbon.Style{
			Fill: carbon.HoverSecondary,
		},
		Active: &carbon.Style{
			Fill: carbon.ActiveSecondary,
		},
		Selected: ButtonSelected,
	}
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
		Selected: ButtonSelected,
	}

	ButtonSelected = &carbon.Style{
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
