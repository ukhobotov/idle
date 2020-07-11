package idle

import (
	"io"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func LoadFont(path string) *truetype.Font {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	var (
		data      []byte
		dataChunk [512]byte
		n         int
	)
	for i := 0; err != io.EOF; i += n {
		n, err = file.ReadAt(dataChunk[:], int64(i))
		data = append(data, dataChunk[:n]...)
	}
	ttf, err := truetype.Parse(data)
	if err != nil {
		panic(err)
	}
	return ttf
}

var (
	Regular *truetype.Font
	Bold    *truetype.Font
)

func init() {
	regular, err := ibmplexsansRegularTtfBytes()
	if err != nil {
		panic(err)
	}
	Regular, err = truetype.Parse(regular)
	if err != nil {
		panic(err)
	}
	bold, err := ibmplexsansBoldTtfBytes()
	if err != nil {
		panic(err)
	}
	Bold, err = truetype.Parse(bold)
	if err != nil {
		panic(err)
	}
}

func NewFace(ttf *truetype.Font, size float64) font.Face {
	face := truetype.NewFace(ttf, &truetype.Options{
		Size: size,
	})
	recover()
	return face
}
