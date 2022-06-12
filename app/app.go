package server

import "net/http"

type App struct {
	mux *http.ServeMux
}

func (a *App) Init() {
	a.mux = http.NewServeMux()
}

func (a *App) NewEndPoint(name string, endpoint func(w http.ResponseWriter, r *http.Request)) {
	a.mux.HandleFunc("/api/v1/"+name, endpoint)
}
