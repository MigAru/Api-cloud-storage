package server

import (
	"net/http"

	structs "github.com/MigAru/Api-cloud-storage/structs"
)


type App struct {
	Mux *http.ServeMux
	Middlewares []func(http.Handler) http.Handler
}

func (a *App) Init(middlewares []func(http.Handler) http.Handler) {
	a.Mux = http.NewServeMux()
	a.Middlewares = middlewares
}

func (a *App) Run(port string) {
	http.ListenAndServe(port, a.Mux)
}

func (a *App) AddEndpoint(pattern string, handler func(http.ResponseWriter, *http.Request)){
	a.Mux.HandleFunc(pattern, handler)
}

func (a *App) RegisterRouterApi(router *structs.Router) {
	for _, endpoint := range router.Endpoints {
		handler := endpoint.Handler
		if len(a.Middlewares) > 0{
			for _, middleware := range a.Middlewares{
				handler = middleware(handler)
			}
		}
		a.Mux.Handle("/api"+router.Version+router.Name+endpoint.Prefix, handler)
	}
}
