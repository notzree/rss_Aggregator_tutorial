package main

import (
	//import json.go file

	"net/http"
)

// implement handler
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithErr(w, 400, "Something went wrong")
}
