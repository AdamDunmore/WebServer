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

func main() {
	checkConfig()
	routes.RegisterRoutes(web, ip, port);
}
