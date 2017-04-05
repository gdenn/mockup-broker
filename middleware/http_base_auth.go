package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// Mockup user credentials
var vc = map[string]string{"username": "John", "password": "test"}

// Use HttpBaseAuth to authenticate user through middleware
func HttpBaseAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authString := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(authString) != 2 || authString[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(authString[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !validate(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Check for valid credentials
func validate(username, password string) bool {
	if vc["username"] == username && vc["password"] == password {
		return true
	}
	return false
}
