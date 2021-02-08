package steganography

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

func stringToBin(s string) string {
	var binString string
	for _, c := range s {
		// 08 is char width. If char has less bits then prepend zeros
		binString = fmt.Sprintf("%s%08b", binString, c)
	}
	return binString
}

func binToString(msg string) string {
	var b []byte

	var prevIDX, nextIDX int
	for nextIDX = 8; nextIDX <= len(msg); nextIDX = nextIDX + 8 {
		sub := msg[prevIDX:nextIDX]
		byteRetrieved, err := strconv.ParseUint(sub, 2, 64)
		if err != nil {
			panic(err)
		}

		b = append(b, byte(byteRetrieved))

		prevIDX = nextIDX
	}

	return string(b)
}

func strPointer(str string) *string {
	return &str
}

func popFirstChar(str *string) string {
	s := *str

	// assign first char
	firstChar := s[:1]

	// pop from string
	*str = s[1:]

	return firstChar
}

func getHash(msg, pin string) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(pin))

	// Write Data to it
	h.Write([]byte(msg))

	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}
