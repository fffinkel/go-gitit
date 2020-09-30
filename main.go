package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
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

func gitInvoke(args ...string) (string, error) {
	actualArgs := make([]string, 2, len(args)+2)
	actualArgs[0] = "-C"
	actualArgs[1] = wikiDir
	actualArgs = append(actualArgs, args...)

	cmd := exec.Command("git", actualArgs...)
	env := []string{}

	for _, envVar := range os.Environ() {
		if strings.HasPrefix(envVar, "GIT_") {
			continue
		}

		if strings.HasPrefix(envVar, "HOME=") {
			continue
		}

		if strings.HasPrefix(envVar, "XDG_CONFIG_HOME=") {
			continue
		}

		env = append(env, envVar)
	}

	env = append(env, "GIT_CONFIG_NOSYSTEM=1")
	env = append(env, "GIT_TRACE=1")

	cmd.Env = env

	output, err := cmd.CombinedOutput()
	fmt.Printf("string(output) = %+v\n", string(output))

	if err != nil {
		return "", err
	}

	return string(output), nil
}

var whiteSpaceRE *regexp.Regexp = regexp.MustCompile(`\s+`)

type noSuchBlobError struct {
	path string
}

func (err *noSuchBlobError) Error() string {
	return "no such blob: " + err.path
}

func (err *noSuchBlobError) StatusCode() int {
	return http.StatusNotFound
}

func (err *noSuchBlobError) Status() string {
	return "Not Found"
}

var _ gititError = &noSuchBlobError{} // make sure noSuchBlobError conforms to gititError interface

func gitGetBlob(path string) (*gitBlob, error) {
	lsTreeOutput, err := gitInvoke("ls-tree", "HEAD", path+".page")
	if err != nil {
		return nil, err // XXX wrap?
	}

	if lsTreeOutput == "" {
		return nil, &noSuchBlobError{path: path}
	}

	blobHash := whiteSpaceRE.Split(strings.Split(strings.TrimSpace(lsTreeOutput), "\n")[0], -1)[2]

	catFileOutput, err := gitInvoke("cat-file", "blob", blobHash)
	if err != nil {
		return nil, err // XXX wrap?
	}

	blob := &gitBlob{
		Hash:     blobHash,
		Contents: catFileOutput,
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
		path = "/Front Page"
	}

	blob, err := gitGetBlob(":" + path)
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

	// playing around!
	http.HandleFunc("/_init", InitHandler())
}

func main() {
	registerHandlers()

	log.Println("Listening on localhost:8001...")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
