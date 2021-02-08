package steganography

import (
	"image"
	"image/color"
	"strconv"
)

// How it works:
// 1. go through each pixel
// 2. for each pixel extract red, green, blue color codes
// 3. transfer them into binary
// 4. change the least significant bit to one of the encodeMessage's bit
// 5. rewrite pixel color with the new one

func encodeMessage(cimg *image.RGBA, binaryString string) {
	encodedPhrase := binaryString

	var rUpdated, gUpdated, bUpdated string
	for i := 0; i < cimg.Bounds().Max.X; i++ {
		for j := 0; j < cimg.Bounds().Max.Y; j++ {
			if encodedPhrase == "" {
				return
			}

			r, g, b, a := cimg.At(i, j).RGBA()
			// get bit representations of colors
			rB := strconv.FormatInt(int64(r), 2)
			gB := strconv.FormatInt(int64(g), 2)
			bB := strconv.FormatInt(int64(b), 2)

			// reassign the least significant bit
			if encodedPhrase != "" {
				rUpdated = rB[:len(rB)-1] + popFirstChar(&encodedPhrase)
			}
			if encodedPhrase != "" {
				gUpdated = rB[:len(gB)-1] + popFirstChar(&encodedPhrase)
			}
			if encodedPhrase != "" {
				bUpdated = rB[:len(bB)-1] + popFirstChar(&encodedPhrase)
			}

			// parse into uint64
			rUpdatedUint64, err := strconv.ParseUint(rUpdated, 2, 64)
			if err != nil {
				panic(err)
			}
			gUpdatedUint64, err := strconv.ParseUint(gUpdated, 2, 64)
			if err != nil {
				panic(err)
			}
			bUpdatedUint64, err := strconv.ParseUint(bUpdated, 2, 64)
			if err != nil {
				panic(err)
			}

			cimg.Set(i, j, color.RGBA{uint8(rUpdatedUint64), uint8(gUpdatedUint64), uint8(bUpdatedUint64), uint8(a)})
		}
	}

	return
}
