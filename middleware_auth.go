package main

import (
	"fmt"
	"net/http"

	"github.com/par1ram/aggregator-go/auth"
	"github.com/par1ram/aggregator-go/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) middleWareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintln("Auth error", err))
		}

		user, err := apiCfg.DB.GetUserByAPIkey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintln("Couldnt get user", err))
		}

		handler(w, r, user)
	}
}
