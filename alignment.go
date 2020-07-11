package idle

type alignment byte

const (
	Left   alignment = 0b1
	Bottom alignment = 0b10
	Right  alignment = 0b100
	Top    alignment = 0b1000
)

func (alt alignment) Has(some alignment) bool {
	return alt&some == some
}
