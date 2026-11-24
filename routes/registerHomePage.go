package routes;

import (
	"net/http"
	"io/fs"
)

func RegisterHomePage(mux *http.ServeMux, files fs.FS, sessions map[string]*bool){
	mux.Handle("/", RegisterAuth(http.FileServer(http.FS(files)), sessions))
	mux.Handle("/style.css", http.FileServer(http.FS(files)))
}


