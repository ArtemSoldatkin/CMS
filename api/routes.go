package api

import (
	"cms/builder"
	"cms/parser"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func makeSite(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	title, children := parser.ReadHTML(body)
	css := []string{"style", "default-style"}
	script := []string{"actions"}
	site := builder.Builder{Title: title, CSS: css, Script: script, Children: children}
	site.Build()
	// return files
}

func createRoutes(r *mux.Router) {
	r.HandleFunc("/make-site", makeSite).Methods(http.MethodPost)
}
