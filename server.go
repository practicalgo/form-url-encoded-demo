package main

import (
	"log"
	"net/http"
)

func main() {
	listenAddr := ":8000"

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/submit-form", handleForm)

	log.Println("starting http server on: ", listenAddr)
	log.Fatal(
		http.ListenAndServe(listenAddr, mux),
	)
}
