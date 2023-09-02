package main

import (
	"log"
	"net/http"
	"shipping-go/handlers/rest"
)

func main() {
	addr := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", rest.TranslateHandler)
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
