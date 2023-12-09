package game

import (
	"github.com/eeyieryi/four-in-a-row/assets/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	SmallFont  font.Face
	MediumFont font.Face
	LargeFont  font.Face
)

func setupFonts() {
	var err error

	tt, err := opentype.Parse(fonts.F_m5x7_ttf)
	if err != nil {
		panic(err)
	}

	SmallFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		panic(err)
	}

	MediumFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		panic(err)
	}

	LargeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    64,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		panic(err)
	}
}
