package structs

import "net/http"

type Endpoint struct {
	Prefix  string
	Handler http.Handler
}

type Router struct {
	Endpoints []Endpoint
	Version string
	Name string
}

func (r *Router) AddNewEndpoint(prefix string, handler func(w http.ResponseWriter, r *http.Request)) {
	endpoint := Endpoint{
		Prefix: prefix,
		Handler: http.HandlerFunc(handler),
	}
	r.Endpoints = append(r.Endpoints, endpoint)
}