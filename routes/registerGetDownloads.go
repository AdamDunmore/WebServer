package routes

import (
	"net/http"
	"strings"
	"os/exec"
)

func RegisterGetDownloads(mux *http.ServeMux, sessions map[string]*bool, music_dir string) {
	mux.Handle("/api/get_downloads", RegisterAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command(
			"ls",
			"-1",  
			music_dir + "/Downloads/",
		) 

		output, err := cmd.CombinedOutput()
		if err != nil {
			http.Error(w, "Getting downloads failed: "+string(output), http.StatusInternalServerError,)
			return
		}

		var results_split []string
		for _, line := range strings.Split(string(output), "\n") {
			if !strings.Contains(line, "Import") && line != "" {
				results_split = append(results_split, line)
			}
		}
		var results string = strings.Join(results_split, "\n"); 
		
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(results))
	}), sessions))

}
