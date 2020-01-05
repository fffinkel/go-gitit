package main

import (
	"log"
	"net/http"
)

func NYI() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "Not yet implemented", http.StatusNotImplemented)
	}
}

func registerHandlers() {
	// Global actions
	http.HandleFunc("/_activity", NYI())
	http.HandleFunc("/_categories", NYI())
	http.HandleFunc("/_index", NYI())
	http.HandleFunc("/_login", NYI())
	http.HandleFunc("/_logout", NYI())
	http.HandleFunc("/_random", NYI())
	http.HandleFunc("/_upload", NYI())
	http.HandleFunc("/_search", NYI())
	http.HandleFunc("/_go", NYI())

	// XXX can take a ?printable param
	http.HandleFunc("/", NYI())

	// Page-specific actions
	http.HandleFunc("/_delete", NYI())
	http.HandleFunc("/_discuss", NYI())
	http.HandleFunc("/_edit", NYI())
	http.HandleFunc("/_history", NYI())
	http.HandleFunc("/_showraw", NYI())

}

func main() {
	registerHandlers()

	log.Println("Listening on localhost:8081...")
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}
