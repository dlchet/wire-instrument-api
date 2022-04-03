package main

import (
	"net/http"

	"github.com/dlchet/wire-instrument-api/server"
)

func main() {
	api := server.SetupAPI()
	router := server.SetupRouter(api)

	http.ListenAndServe(":3000", router)
}
