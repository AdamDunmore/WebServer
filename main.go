package main

import (
	"embed"
	"encoding/json"
	"os"
	"io/fs"
	"log"
	"net/http"
	"os/exec"
	"crypto/rand"
	"encoding/hex"
	"net/url"
)

//go:embed web/*
var web embed.FS

var sessions = make(map[string]bool)

func validSession(id string) bool {
	return sessions[id]
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil || !validSession(cookie.Value) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type DownloadRequest struct {
	Name string `json:"name"`
}

type LoginRequest struct {
	Password string `json:"password"`
}

func main() {
	// ip := "100.99.196.79"
	ip := "localhost"
	port := "1913"

	files, err := fs.Sub(web, "web")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// Login page -- currently public
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, files, "login.html")
	})

	// Download API -- currently protected
	mux.Handle("/api/download", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data DownloadRequest

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		cmdURL, err := url.ParseRequestURI(data.Name)
		if err != nil || (cmdURL.Scheme != "http" && cmdURL.Scheme != "https") {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
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
			"url",
			data.Name,
		)

		if err := cmd.Run(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Download " + data.Name + " successful"))
	})))

	// Login API -- currently unprotected
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
		secure := ip != "localhost" && ip != "127.0.0.1"
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    sessionID,
			HttpOnly: true,
			Secure:   secure,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		})

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	}))

	// Website -- protected
	mux.Handle("/", auth(
		http.FileServer(http.FS(files)),
	))

	mux.Handle("/style.css", http.FileServer(http.FS(files)))

	log.Println("Listening on " + ip + ":" + port)
	log.Fatal(http.ListenAndServe(ip+":"+port, mux))
}
