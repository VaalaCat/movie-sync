package main

import (
	"embed"
	"movie-sync-server/services"
)

//go:embed all:out
var fs embed.FS

func main() {
	services.EventHandler()
	services.RouterHandler(fs)
	services.Run()
}
