package main

import (
	"net/http"

	server "github.com/MigAru/Api-cloud-storage/app"
)

func main() {
	app := server.App{}
	middlewares := []func(http.Handler) http.Handler{}
	app.Init(middlewares)
}
