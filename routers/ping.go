package routers

import (
	"net/http"

	structs "github.com/MigAru/Api-cloud-storage/structs"
)

var Pingrouter structs.Router = structs.Router{
	Version: "/v1",
	Name: "/ping",
	Endpoints: []structs.Endpoint{
		pingEndpoint,
	},
	}

var pingEndpoint structs.Endpoint = structs.Endpoint{
	Prefix: "/",
	Handler: http.HandlerFunc(Ping),

}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"response": "pong"}`))
}