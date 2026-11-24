package routes;

import (
	"net/http"
	"io/fs"
)

func RegisterLoginPage(mux *http.ServeMux, files fs.FS){
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, files, "login.html")
	})
}


