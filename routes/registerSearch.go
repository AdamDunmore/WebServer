package routes;

import (
	"encoding/json"
	"os"
	"net/http"
	"os/exec"
)

import "webserver/routes/classes"

func RegisterSearch(mux *http.ServeMux, sessions map[string]*bool){
	mux.Handle("/api/search", RegisterAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data classes.MediaRequest

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		cmd := exec.Command(
			"rip",
			"search",
			"qobuz",
			"-o",
			"/tmp/rip-results",
			data.MediaType,
			data.Id,
		)

		output, err := cmd.CombinedOutput()
		if err != nil {
			http.Error(w, "rip failed: "+string(output), http.StatusInternalServerError,)
			return
		}

		fileData, err := os.ReadFile("/tmp/rip-results")
		if err != nil {
			http.Error(w, "Failed to read results", http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.Write(fileData)
	}), sessions))
}
