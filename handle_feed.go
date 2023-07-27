package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"time"
)

// implement handler
func (h *Handler) handleCreateFeed(w http.ResponseWriter, r *http.Request) {
	// type parameters struct {
	// 	Name string `json:"name"`
	// }
	var params Feeds
	err := json.NewDecoder(r.Body).Decode(&params)
	log.Println(params)
	if params.UserId == 0 {
		respondWithErr(w, 400, fmt.Sprintf("Missing user id ", params.UserId))
		return
	}

	params.Id, _ = NewId()
	params.Created_at = time.Now()
	params.Updated_at = time.Now()
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing JSON", err))
		return
	}
	response := h.db.Create(&params)
	if response.Error != nil {
		respondWithErr(w, 500, fmt.Sprintf("Error creating feed", response.Error))
		return
	}
	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(params))
}
func (h *Handler) handleReadFeed(w http.ResponseWriter, r *http.Request) {
	var params QueryFeed
	var feed Feeds
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error parsing JSON", err))
		return
	}

	response := h.db.First(&feed, params.Id)
	if response.Error != nil {
		respondWithErr(w, 500, fmt.Sprintf("Error reading feed", response.Error))
		return
	}
	log.Println(feed)
	respondWithJSON(w, http.StatusOK, feed)
}
