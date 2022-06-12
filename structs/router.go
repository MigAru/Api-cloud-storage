package structs

import "net/http"

type Endpoint struct {
	prefix  string
	handler func(w http.ResponseWriter, r *http.Request)
}

type Router struct {
	endpoints []Endpoint
	version string
	name string
}
