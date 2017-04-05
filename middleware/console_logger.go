package middleware

import "net/http"

// Log requests to console
func ConsoleLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log to console
	})
}
