package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/par1ram/aggregator-go/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params) // unMarshalling JSON to parrams
	if err != nil {
		respondWithError(w, 400, fmt.Sprintln("Error parsing JSON", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintln("Couldnt create user", err))
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feed, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintln("Couldnt get feeds", err))
	}

	respondWithJSON(w, 201, databaseFeedsToFeeds(feed))
}
