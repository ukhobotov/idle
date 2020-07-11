package styles

import (
	"github.com/faiface/pixel"
	"github.com/fogleman/gg"
	"github.com/usualhuman/idle"
	"github.com/usualhuman/idle/carbon"
)

type Input struct {
	idle.Location
	Style idle.HandlerStyle
}

func (input *Input) Rasterize() {
	panic("implement me")
}

func (input *Input) Handle(event idle.Event) {
	panic("implement me")
}

func (input *Input) Draw(target pixel.Target) {

}

func TextInput1(loc idle.Location, placeholder string) idle.Element {
	var input string
	return &idle.Container{
		Location: loc,
		Content: []idle.Element{
			&idle.Handler{
				Style: idle.HandlerStyle{
					Idle: &idle.State{
						Fill: carbon.Field1,
						Border: idle.Border{
							Color:     carbon.Ui4,
							Alignment: idle.Bottom,
						},
					},
				},
				Final: func(ctx *gg.Context) {
					ctx.SetFontFace(idle.NewFace(idle.Regular, 16))
					if input == "" {
						ctx.SetColor(carbon.Text3)
					}
				},
				OnPress: func(this *idle.Handler) {

				},
			},
		},
	}
}
