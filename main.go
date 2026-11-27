package main

import (
	"embed"

	"webserver/routes"
)

// Create embed
//go:embed web/*
var web embed.FS

// Config values
var ip string
var port string
var password string
var music_dir string

func main() {
	globals := checkConfig()
	routes.RegisterRoutes(web, globals);
}
