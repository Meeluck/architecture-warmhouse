package main

import (
	"log"
	"net/http"

	"github.com/Meeluck/architecture-warmhouse/temperature-api/api"
)

func main() {
	http.HandleFunc("/temperature", api.TemperatureHandler)
	addr := ":8081"
	log.Printf("temperature-api started on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
