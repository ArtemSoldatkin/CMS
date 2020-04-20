package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateAPI - create api
func CreateAPI() {
	r := mux.NewRouter()
	createRoutes(r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
