package app

import (
	"net/http"

	u "go-learn/rest2/utils"
)

var jwtAuth = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		skipAuth := map[string]bool{
			"api/user/new":   true,
			"api/user/login": true,
		}
		if skipAuth[r.URL.Path] {
			next.ServeHTTP(w, r)
			return
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}
	})
}
