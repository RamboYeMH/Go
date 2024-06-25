package middleware

import (
	"log"
	"mime"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("recv a %s request from %s", request.Method, request.RemoteAddr)
		next.ServeHTTP(writer, request)
	})
}

func Validating(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		contentType := request.Header.Get("Content-Type")
		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		if mediaType != "application/json" {
			http.Error(writer, "invalid Content-Type", http.StatusUnsupportedMediaType)
			return
		}
		next.ServeHTTP(writer, request)
	})
}
