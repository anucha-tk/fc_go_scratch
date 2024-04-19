package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/anucha-tk/fc_go_scratch/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Url:       params.URL,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error can't create feed: %v", err))
		return
	}
	responseWithJSON(w, 201, feed)
}

// TODO: make pagination
func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request, _ database.User) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error can't get feeds: %v", err))
		return
	}

	responseWithJSON(w, 200, feeds)
}
