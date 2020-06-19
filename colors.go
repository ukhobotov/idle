package carbon

import (
	"fmt"
	"image/color"
)

var (
	Black       = color.RGBA{0, 0, 0, 255}
	White       = color.RGBA{255, 255, 255, 255}
	Transparent = color.RGBA{}

	Red70 = hex("#a2191f")
	Red60 = hex("#da1e28")
	Red50 = hex("#fa4d56")
	Red40 = hex("#ff8389")
	Red30 = hex("#ffb3b8")
	Red20 = hex("#ffd7d9")
	Red10 = hex("#fff1f1")

	Orange50 = hex("#eb6200")
	Orange40 = hex("#ff832b")
	Orange30 = hex("#ffb784")
	Orange20 = hex("#ffd9be")

	Yellow50 = hex("#b28600")
	Yellow40 = hex("#d2a106")
	Yellow30 = hex("#f1c21b")
	Yellow20 = hex("#fddc69")
	Yellow10 = hex("#fcf4d6")

	Green80 = hex("#044317")
	Green70 = hex("#0e6027")
	Green60 = hex("#198038")
	Green50 = hex("#24a148")
	Green40 = hex("#42be65")
	Green30 = hex("#6fdc8c")
	Green20 = hex("#a7f0ba")

	Teal50 = hex("#009d9a")
	Teal40 = hex("#08bdba")
	Teal30 = hex("#3ddbd9")
	Teal20 = hex("#9ef0f0")

	Cyan50 = hex("#1192e8")
	Cyan40 = hex("#33b1ff")
	Cyan30 = hex("#82cfff")
	Cyan20 = hex("#bae6ff")

	Blue100 = hex("#001141")
	Blue90  = hex("#001d6c")
	Blue80  = hex("#002d9c")
	Blue70  = hex("#0043ce")
	Blue60  = hex("#0f62fe")
	Blue50  = hex("#4589ff")
	Blue40  = hex("#78a9ff")
	Blue30  = hex("#a6c8ff")
	Blue20  = hex("#d0e2ff")
	Blue10  = hex("#edf5ff")

	Purple60 = hex("#8a3ffc")
	Purple50 = hex("#a56eff")
	Purple40 = hex("#be95ff")
	Purple30 = hex("#d4bbff")

	Magenta60 = hex("#d12771")
	Magenta50 = hex("#ee5396")
	Magenta40 = hex("#ff7eb6")
	Magenta30 = hex("#ffafd2")

	Gray100 = hex("#161616") // Balanced
	Gray90  = hex("#262626")
	Gray80  = hex("#393939")
	Gray70  = hex("#525252")
	Gray60  = hex("#6f6f6f")
	Gray50  = hex("#8d8d8d")
	Gray30  = hex("#c6c6c6")
	Gray10  = hex("#f4f4f4")

	// Gray100 = hex("#121619") // Cooler ones
	// Gray90  = hex("#21272a")
	// Gray80  = hex("#343a3f")
	// Gray70  = hex("#4d5358")
	// Gray60  = hex("#697077")

	Blue60Hover = hex("#0353e9")
	Gray90Hover = hex("#353535")
	Gray60Hover = hex("#606060")

	// Default Background
	Bg = Gray100

	// ButtonPrimary interactive color; ButtonPrimary button
	Interactive1 = Blue60
	// Secondary interactive color; Secondary button
	Interactive2 = Gray60
	// Tertiary button
	Interactive3 = White
	// Selected elements; Activated elements; Accent tools
	Interactive4 = Blue50

	// Container background on Bg; Secondary page background
	Ui1 = Gray90
	// Container background on Ui1; `Light` variant Background
	Ui2 = Gray80
	// Subtle border; Tertiary background
	Ui3 = Gray80
	// Medium contrast border
	Ui4 = Gray60
	// High contrast border; Emphasis elements
	Ui5 = Gray10

	// ButtonPrimary text
	Text1 = Gray10
	// Secondary text
	Text2 = Gray30
	// Placeholder text
	Text3 = Gray60
	// Text on interactive colors
	Text4 = White
	// Tertiary text
	Text5 = Gray50

	// ButtonPrimary icons
	Icon1 = Gray10
	// Secondary icons
	Icon2 = Gray30
	// Icons on interactive colors; Icons on non-ui colors
	Icon3 = White

	// Interactive1 hover
	HoverPrimary = Blue60Hover
	// Interactive1 text hover
	HoverPrimaryText = Blue70
	// Interactive2 hover
	HoverSecondary = Gray60Hover
	// Interactive3 hover; Inverse1 hover
	HoverTertiary = Gray10
	// Ui1 hover; Ui2 hover; Transparent background hover
	HoverUi = Gray90Hover

	// Interactive1 active
	ActivePrimary = Blue80
	// Interactive2 active; Inverse1 active
	ActiveSecondary = Gray80
	// Interactive3 active
	ActiveTertiary = Gray30
	// Ui1 active; Ui2 active
	ActiveUi = Gray70
)

func hex(hex string) color.RGBA {
	if hex[0] != '#' || len(hex) != 7 {
		panic(fmt.Errorf("wrong color format: %s", hex))
	}
	var bytes [6]byte
	for i, r := range hex[1:] {
		b := byte(r)
		switch {
		case '0' <= b && b <= '9':
			bytes[i] = b - '0'
		case 'a' <= b && b <= 'f':
			bytes[i] = 10 + b - 'a'
		default:
			panic(fmt.Errorf("wrong color format: %s", hex))
		}
	}
	return color.RGBA{
		R: bytes[0]<<4 + bytes[1],
		G: bytes[2]<<4 + bytes[3],
		B: bytes[4]<<4 + bytes[5],
		A: 0xff,
	}
}
