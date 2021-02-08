package steganography

import (
	"net/http"
	"strings"
)

type cors struct {
	allowedOrigins string
	allowedMethods string
	allowedHeaders string
}

var (
	corsAcceptedHeaders = []string{
		"Accept",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"X-Apollo-Tracing",
	}
	corsAllowedMethods = []string{"POST", "GET", "OPTIONS", "PUT", "PATCH", "DELETE"}
)

// HandleCORS sets response headers and handles preflight requests
func HandleCORS(h http.HandlerFunc, options *cors) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//set headers
		if options != nil && options.allowedOrigins != "" {
			w.Header().Set("Access-Control-Allow-Origin", options.allowedOrigins)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		if options != nil && options.allowedMethods != "" {
			w.Header().Set("Access-Control-Allow-Methods", options.allowedMethods)
		} else {
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(corsAllowedMethods, ","))
		}

		if options != nil && options.allowedHeaders != "" {
			w.Header().Set("Access-Control-Allow-Headers", options.allowedHeaders)
		} else {
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(corsAcceptedHeaders, ","))
		}

		//handle pre-flight requests
		if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}
