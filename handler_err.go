package main

import (
	"net/http"
)

// implement handler
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithErr(w, 400, "Something went wrong")
}
