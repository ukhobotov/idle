package carbon

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

func defaultFont() *truetype.Font {
	// ttf, _ := truetype.Parse() TODO: fix
	return nil
}

var (
	Regular = LoadFont("IBMPlexSans-Regular.ttf")
	Bold    = LoadFont("IBMPlexSans-Bold.ttf")
)

func NewFace(ttf *truetype.Font, size float64) font.Face {
	face := truetype.NewFace(ttf, &truetype.Options{
		Size: size,
	})
	recover()
	return face
}
