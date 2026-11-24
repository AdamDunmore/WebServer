package routes;

import (
	"net/http"
)

func RegisterAuth(next http.Handler, sessions map[string]*bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if ((err != nil) || (sessions[cookie.Value] == nil)){
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
