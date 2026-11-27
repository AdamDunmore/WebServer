package routes;

import (
	"encoding/json"
	"net/http"
)

import "webserver/routes/classes"

func RegisterLogin(mux *http.ServeMux, sessions map[string]*bool, password string){
	mux.Handle("/api/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data classes.LoginRequest

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if data.Password != password {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		var status bool = AttemptLogin(r, w, sessions);
		if (!status) {
			http.Error(w, "Failed to create session", http.StatusInternalServerError);
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	}))
}
