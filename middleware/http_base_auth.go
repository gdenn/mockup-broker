package middleware

import "net/http"

const validUserCredentials := make(map[string]string) {
"name": "John",
"password"
}

// Use HttpBaseAuth to authenticate user through middleware
func HttpBaseAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !authenticate(r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func authenticate(r *http.Request) bool {
	return true
}
