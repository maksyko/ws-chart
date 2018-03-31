package main

import (
	"github.com/gorilla/mux"
	"github.com/ievgen-ma/ws-chart/messaging"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	messaging.Start(r)
	http.ListenAndServe(":3000", r)
}
