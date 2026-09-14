package main

import (
	"encoding/json"
	"os"
	"io/fs"
	"net/http"
	"os/exec"
	"crypto/rand"
	"encoding/hex"
)

var sessions = make(map[string]bool)

type LoginRequest struct {
	Password string `json:"password"`
}

type MediaRequest struct {
	Id string `json:"id"`
	MediaType string `json:"media_type"`
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil || !sessions[cookie.Value] {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func registerSearch(mux *http.ServeMux){
	mux.Handle("/api/search", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data MediaRequest

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
	})))
}

func registerDownloadId(mux *http.ServeMux){
	mux.Handle("/api/download_id", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data MediaRequest

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
	})))
}

func registerLogin(mux *http.ServeMux, files fs.FS){
	// Login page -- currently public
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, files, "login.html")
	})
}

func registerApiLogin(mux *http.ServeMux){
	mux.Handle("/api/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data LoginRequest

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if data.Password != os.Getenv("WEBSERVER_PASSWORD") {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		// Generate random session ID
		b := make([]byte, 32)
		if _, err := rand.Read(b); err != nil {
			http.Error(w, "Failed to create session", http.StatusInternalServerError)
			return
		}

		sessionID := hex.EncodeToString(b)

		// Mark session as authenticated
		sessions[sessionID] = true

		// Give session to browser
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    sessionID,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		})

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	}))
}

func registerWeb(mux *http.ServeMux, files fs.FS){
	mux.Handle("/", auth(http.FileServer(http.FS(files))))
	mux.Handle("/style.css", http.FileServer(http.FS(files)))
}
