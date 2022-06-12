package server

import (
	"net/http"

	structs "github.com/MigAru/Api-loud-storage/structs"
)


type App struct {
	mux *http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

func (a *App) Init(middlewares []func(http.Handler) http.Handler) {
	a.mux = http.NewServeMux()
}

func (a *App) Run(port string) {
	http.ListenAndServe(port, a.mux)
}
func (a *App) RegisterRouterApi(router *structs.Router) {
	for _, endpoint := range router.Endpoints {
		handler := endpoint.Handler
		if len(a.middlewares) > 0{
			for _, middleware := range a.middlewares{
				handler = middleware(handler)
			}
		}
		a.mux.Handle("/api"+router.Version+router.Name+endpoint.Prefix, handler)
	}
}
