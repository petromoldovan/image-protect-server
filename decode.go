package steganography

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"strconv"
)

func retrieveMessage(img *image.RGBA) (string, error) {
	var msg string
	for i := 0; i < img.Bounds().Max.X; i++ {
		for j := 0; j < img.Bounds().Max.Y; j++ {
			if m, err := checkForEnd(msg); err != nil {
				return m, nil
			}

			r, g, b, _ := img.At(i, j).RGBA()

			// get bit representations of colors
			rB := strconv.FormatInt(int64(r), 2)
			gB := strconv.FormatInt(int64(g), 2)
			bB := strconv.FormatInt(int64(b), 2)

			if m, err := checkForEnd(msg); err != nil {
				return m, nil
			}
			msg = fmt.Sprintf("%s%s", msg, rB[len(rB)-1:])

			if m, err := checkForEnd(msg); err != nil {
				return m, nil
			}
			msg = fmt.Sprintf("%s%s", msg, gB[len(gB)-1:])

			if m, err := checkForEnd(msg); err != nil {
				return m, nil
			}
			msg = fmt.Sprintf("%s%s", msg, bB[len(bB)-1:])
		}
	}

	return "", fmt.Errorf("could not retrieve the message")
}

//-------------
// Helpers

func checkForEnd(msg string) (string, error) {
	// msg length must be divisible by bit length of a char
	if len(msg)%8 != 0 {
		return "", nil
	}

	if len(msg) >= len(ENCODING_END) && msg[len(msg)-len(ENCODING_END):] == ENCODING_END {
		return msg[:len(msg)-len(ENCODING_END)], fmt.Errorf("reached end")
	}
	return "", nil
}
