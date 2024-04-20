package middleware

import (
	"log"
	"net/http"
)

// now its unsafe becouse of using only simple secret on the both sides
// we must make some auth functions that can compare its result with result of another auth function
func Encryptor(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if secret != "" {
				hashHeader := r.Header.Get("HashSHA256")
				if hashHeader != secret {
					log.Print("YOU SHALL NOT PASS!")
					http.Error(w, "unvalid secret", http.StatusForbidden)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
