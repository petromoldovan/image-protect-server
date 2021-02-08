package main

import (
	"fmt"
	"log"
	"os"

	"github.com/petromoldovan/image-protect-server"
)

func main() {
	steganography.Log(fmt.Sprintf("Starting server on port: %s", steganography.SERVER_PORT))
	if err := steganography.StartServer(); err != nil {
		log.Fatal("ListenAndServe: ", err)
		os.Exit(1)
	}
}
