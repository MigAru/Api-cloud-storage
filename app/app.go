package server

import (
	"net/http"

	structs "github.com/MigAru/Api-loud-storage/structs"
)


type App struct {
	mux *http.ServeMux
}

func (a *App) Init() {
	a.mux = http.NewServeMux()
}

func (a *App) Run(port string) {
	http.ListenAndServe(port, a.mux)
}

func (a *App) RegisterRouter(router *structs.Router) {
	
}
