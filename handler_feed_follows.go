package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/anucha-tk/fc_go_scratch/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error can't create feed_follow: %v", err))
		return
	}
	responseWithJSON(w, 201, feedFollow)
}

func (apiCfg *apiConfig) handlerGetFeedFollowByUserID(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollowsByUserId(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, 404, fmt.Sprintf("Error can't find feedFollows: %v", err))
		return
	}
	responseWithJSON(w, 200, feedFollows)
}

func (apiCfg *apiConfig) handlerDeleteFeedFollowByID(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err == nil {
		responseWithError(w, 400, fmt.Sprintf("invalid uuid: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollowById(r.Context(), database.DeleteFeedFollowByIdParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error can't delete feedFollowById: %v", err))
		return
	}
	responseWithJSON(w, 200, struct{}{})
}
