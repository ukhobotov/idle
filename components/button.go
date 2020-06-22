package components

import "carbon"

type (
	Button struct {
		handler
		Style ButtonStyle
		Text  *ButtonText
		Icon  *ButtonIcon
	}

	ButtonStyle byte
	ButtonText  struct {
		String string
		X, Y   float64
	}
	ButtonIcon struct {
		Drawing carbon.Drawing
		X, Y    float64
	}
)

const (
	Primary ButtonStyle = iota
	Secondary
	Tertiary
)

// func Button(loc carbon.Location, text string, textX, textY, text)
