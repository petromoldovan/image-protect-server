package steganography

import (
	"log"
	"os"
	"time"
)

func Log(msg string) {
	l := log.New(os.Stdout, "", 0)
	l.SetPrefix(time.Now().UTC().Format("2006-01-02 15:04:05") + " ")
	l.Print(msg)
}
