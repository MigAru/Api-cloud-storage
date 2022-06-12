package main

import (
	"fmt"
	"net/http"

	server "github.com/MigAru/Api-loud-storage/app"
)

func main() {
	app := server.App{}
	middlewares := []func(http.Handler) http.Handler{}
	app.Init(middlewares)
	fmt.Println(len(middlewares))
}
