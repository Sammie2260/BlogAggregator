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

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondwithJSON(w, 201, databaseUserToUser(user))

}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// remove user database.user if uncommented.
	// apiKey, err := auth.GetAPIKey(r.Header)
	// if err != nil {
	// 	respondwithError(w, 403, fmt.Sprintf("Auth error: %v", err))
	// 	return
	// }
	// user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	// if err != nil {
	// 	respondwithError(w, 403, fmt.Sprintf("Couldn't get user: %v", err))
	// 	return
	// }
	respondwithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostForUser(r.Context(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	respondwithJSON(w, 200, databasePostsToPosts(posts))
}
