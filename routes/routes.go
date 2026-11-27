package routes

import (
	"crypto/rand"
	"embed"
	"encoding/hex"
	"io/fs"
	"log"
	"net/http"
)

func CreateSessions() map[string]*bool {
	var sessions = make(map[string]*bool)
	return sessions
}

func makeSessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		println("Failed to create session ID")
		return "null";
	}
	return hex.EncodeToString(b);	
}

func AttemptLogin(r *http.Request, w http.ResponseWriter, sessions map[string]*bool) bool {
	cookie, err := r.Cookie("session")
	var session_id string;
	if (err != nil || cookie.Value == "null") { 
		session_id = makeSessionID();
	} else {
		session_id = cookie.Value; 
	}	
	
	val := true;
	sessions[session_id] = &val;

	// Give session to browser
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    session_id,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
	
	return true;	
}

func RegisterRoutes(web embed.FS, globals[] *string){
	var password string = *globals[0];
	var ip string = *globals[1];
	var port string = *globals[2];
	var music_dir = *globals[3];

	// Gets /web files
	files, err := fs.Sub(web, "web")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux();
	sessions := CreateSessions();
	RegisterHomePage(mux, files, sessions);
	RegisterLoginPage(mux, files);
	RegisterLogin(mux, sessions, password);
	RegisterSearch(mux, sessions);
	RegisterDownloadId(mux, sessions, music_dir);
	RegisterGetDownloads(mux, sessions, music_dir);

	// Print status to log
	log.Println("Listening on " + ip + ":" + port)
	log.Fatal(http.ListenAndServe(ip + ":" + port, mux))
}
