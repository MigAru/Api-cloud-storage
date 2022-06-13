package main

import (
	"net/http"

	server "github.com/MigAru/Api-cloud-storage/app"
	middlewares "github.com/MigAru/Api-cloud-storage/middlewares"
	route "github.com/MigAru/Api-cloud-storage/routers"
)

func main() {
	App := server.App{}
	middlewares := []func(http.Handler) http.Handler{middlewares.Logger,}
	App.Init(middlewares)
	App.RegisterRouterApi(&route.Pingrouter)
	App.Run("8080")
}
