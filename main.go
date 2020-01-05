package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func NYI() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "Not yet implemented", http.StatusNotImplemented)
	}
}

type gititError interface {
	error

	StatusCode() int
	Status() string
}

type gitBlob struct {
	Hash     string
	Contents string
}

func gitGetBlob(path string) (*gitBlob, error) {
	blob := &gitBlob{
		Hash:     "1d9616855a130da2cd0665f79139f6d7853595b1", // XXX STUB
		Contents: "lorem ipsum\n",
	}

	return blob, nil
}

// XXX STUB
func renderMarkdown(markdown string) (string, error) {
	return markdown, nil
}

// XXX can take a ?printable param
func rootHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if path == "/" {
		path = "Front Page"
	}

	blob, err := gitGetBlob(path)
	if err != nil {
		if appErr, ok := err.(gititError); ok {
			// XXX write out an HTML page that shows the error message?
			http.Error(w, appErr.Status(), appErr.StatusCode())
		} else {
			log.Println(err.Error()) // XXX show this on a page in debug mode or something?
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Add("ETag", blob.Hash)
	renderedMarkdown, err := renderMarkdown(blob.Contents)
	if err != nil {
		log.Println(err.Error()) // XXX show this on a page in debug mode or something?
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.ServeContent(w, req, "", time.Now(), strings.NewReader(renderedMarkdown))
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

	http.HandleFunc("/", rootHandler)

	// Page-specific actions
	http.HandleFunc("/_delete", NYI())
	http.HandleFunc("/_discuss", NYI())
	http.HandleFunc("/_edit", NYI())
	http.HandleFunc("/_history", NYI())
	http.HandleFunc("/_showraw", NYI())

}

func main() {
	registerHandlers()

	log.Println("Listening on localhost:8001...")
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}
