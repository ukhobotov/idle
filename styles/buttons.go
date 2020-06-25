package styles

import (
	"github.com/fogleman/gg"

	"github.com/usualhuman/carbon"
)

// Button are elements user primarily for actions. There are some kinds of it that might be used for special purposes.

// PrimaryButton is needed for the principal call to action on the page.
func PrimaryButton() carbon.HandlerStyle {
	return carbon.HandlerStyle{
		Idle: &carbon.Style{
			Fill: carbon.Interactive1,
		},
		Hover: &carbon.Style{
			Fill: carbon.HoverPrimary,
		},
		Active: &carbon.Style{
			Fill: carbon.ActivePrimary,
		},
		Focus: ButtonFocus(),
		Common: func(ctx *gg.Context) {
			ctx.SetColor(carbon.Text4)
		},
	}
}

// SecondaryButton is needed for secondary actions on each page.
func SecondaryButton() carbon.HandlerStyle {
	return carbon.HandlerStyle{
		Idle: &carbon.Style{
			Fill: carbon.Interactive2,
		},
		Hover: &carbon.Style{
			Fill: carbon.HoverSecondary,
		},
		Active: &carbon.Style{
			Fill: carbon.ActiveSecondary,
		},
		Focus: ButtonFocus(),
		Common: func(ctx *gg.Context) {
			ctx.SetColor(carbon.Text4)
		},
	}
}

// TertiaryButton is needed for some additional actions.
func TertiaryButton() carbon.HandlerStyle {
	return carbon.HandlerStyle{
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
		Focus: ButtonFocus(),
	}
}

func ButtonFocus() *carbon.Style {
	return &carbon.Style{
		Border: carbon.Border{
			Color:    carbon.Focus,
			Width:    2,
			Splitter: carbon.Background,
		},
	}
}
