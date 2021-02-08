package steganography

import (
	"fmt"
	"net/http"
)

func StartServer() error {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/encode", HandleCORS(encodeHandler, nil))
	http.HandleFunc("/decode", HandleCORS(decodeHandler, nil))

	return http.ListenAndServe(fmt.Sprintf("%s:%s", SERVER_URL, SERVER_PORT), nil)
}
