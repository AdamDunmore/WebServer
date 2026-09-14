package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed web/*
var web embed.FS

var ip string
var port string
var password string

func main() {
	checkConfig()

	// Gets /web files
	files, err := fs.Sub(web, "web")
	if err != nil {
		log.Fatal(err)
	}

	// Sets up router
	mux := http.NewServeMux()

	// Register Routes
	registerDownloadId(mux)
	registerSearch(mux)
	registerLogin(mux, files)
	registerApiLogin(mux)
	registerWeb(mux, files)

	// Print status to log
	log.Println("Listening on " + ip + ":" + port)
	log.Fatal(http.ListenAndServe(ip + ":" + port, mux))
}
