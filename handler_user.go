package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// implement handler
func (h *Handler) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	var params Users
	err := json.NewDecoder(r.Body).Decode(&params)
	params.Id, _ = NewId()
	params.Created_at = time.Now()
	params.Updated_at = time.Now()
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing JSON", err))
		return
	}
	log.Println(params)
	response := h.db.Create(&params)
	if response.Error != nil {
		respondWithErr(w, 500, fmt.Sprintf("Error creating user", response.Error))
		return
	}
	respondWithJSON(w, http.StatusOK, databaseUserToUser(params))
}
func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON body from the request
	var user Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithErr(w, http.StatusBadRequest, "Error parsing JSON")
		return
	}
	//Preload loads the associated fields, ie ones that are related to the user type.
	// In this case, it loads the feeds
	h.db.Preload("Feeds").Find(&user, user.Id)

	// Respond with the feeds from the specific user
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
