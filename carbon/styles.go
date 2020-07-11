package carbon

import (
	"github.com/usualhuman/idle"
	"github.com/usualhuman/idle/carbon/components"
)

// PrimaryButton is needed for the principal call to action on the page.
func PrimaryButton() *components.ButtonStyle {
	return &components.ButtonStyle{
		Idle: &idle.State{
			Fill: Interactive1,
			Drawer: &idle.TextIconDrawer{
				TextColor: Text4,
				IconColor: Icon3,
			},
		},
		Hover: &idle.State{
			Fill: HoverPrimary,
			Drawer: &idle.TextIconDrawer{
				TextColor: Text4,
				IconColor: Icon3,
			},
		},
		Active: &idle.State{
			Fill: ActivePrimary,
			Drawer: &idle.TextIconDrawer{
				TextColor: Text4,
				IconColor: Icon3,
			},
		},
		Focus: &idle.State{
			Border: idle.Border{
				Color: Focus,
				Width: 2,
				Inset: Inverse1,
			},
		},
		Disabled: &idle.State{
			Fill: Disabled2,
			Drawer: &idle.TextIconDrawer{
				TextColor: Disabled3,
				IconColor: Disabled3,
			},
		},
	}
}

// SecondaryButton is needed for secondary actions on each page.
func SecondaryButton() *components.ButtonStyle {
	return &components.ButtonStyle{
		Idle: &idle.State{
			Fill: Interactive2,
			Drawer: &idle.TextIconDrawer{
				TextColor: Text4,
				IconColor: Icon3,
			},
		},
		Hover: &idle.State{
			Fill: HoverSecondary,
			Drawer: &idle.TextIconDrawer{
				TextColor: Text4,
				IconColor: Icon3,
			},
		},
		Active: &idle.State{
			Fill: ActiveSecondary,
			Drawer: &idle.TextIconDrawer{
				TextColor: Text4,
				IconColor: Icon3,
			},
		},
		Focus: &idle.State{
			Border: idle.Border{
				Color: Focus,
				Width: 2,
				Inset: Inverse1,
			},
		},
		Disabled: &idle.State{
			Fill: Disabled2,
			Drawer: &idle.TextIconDrawer{
				TextColor: Disabled3,
				IconColor: Disabled3,
			},
		},
	}
}

// TertiaryButton is needed for some additional actions.
func TertiaryButton() *components.ButtonStyle {
	return &components.ButtonStyle{
		Idle: &idle.State{
			Border: idle.Border{
				Color: Interactive3,
			},
		},
		Hover: &idle.State{
			Fill: HoverTertiary,
		},
		Active: &idle.State{
			Fill: ActiveTertiary,
		},
	}
}
