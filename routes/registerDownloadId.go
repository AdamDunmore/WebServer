package routes;

import (
	"encoding/json"
	"os"
	"net/http"
	"os/exec"
)

import "webserver/routes/classes"

func RegisterDownloadId(mux *http.ServeMux, sessions map[string]*bool){
	mux.Handle("/api/download_id", RegisterAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data classes.MediaRequest

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		var home = os.Getenv("HOME");
		cmd := exec.Command(
			"rip",
			"-q",
			"3",
			"-f",
			home + "/Music/Downloads/",
			"-c",
			"FLAC",
			"id",
			"qobuz",
			data.MediaType, // TODO artist (downloads ALL music), playlist (maybe change download dir?) 
			data.Id,
		)
		
		if err := cmd.Run(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Download " + data.Id + " successful"))
	}), sessions))
}


