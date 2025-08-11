package main

import (
	"encoding/json"
	"fmt"

	//"hellogo/internal/auth"
	"hellogo/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	respondwithJSON(w, 201, databaseFeedtoFeed(feed))

}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	respondwithJSON(w, 201, databaseFeedstoFeeds(feeds))

}
