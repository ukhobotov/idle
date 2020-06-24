package styles

import (
	"image/color"

	"github.com/fogleman/gg"

	"github.com/usualhuman/carbon"
)

// Button are elements user primarily for actions. There are some kinds of it that might be used for special purposes.

var (
	// PrimaryButton is needed for the principal call to action on the page.
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
		Focus: ButtonFocus,
		Common: func(ctx *gg.Context) {
			ctx.SetColor(carbon.Text4)
		},
	}

	// SecondaryButton is needed for secondary actions on each page.
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
		Focus: ButtonFocus,
		Common: func(ctx *gg.Context) {
			ctx.SetColor(carbon.Text4)
		},
	}

	// TertiaryButton is needed for some additional actions.
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
		Focus: ButtonFocus,
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
