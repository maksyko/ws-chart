package main

import (
	"github.com/gorilla/mux"
	"github.com/ws-chart/messaging"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	messaging.Start(r)
	http.ListenAndServe(":3000", r)
}
