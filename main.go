package main

import "movie-sync-server/services"

func main() {
	services.EventHandler()
	services.RouterHandler()
	services.Run()
}
