package main

import (
	"fmt"
	"net/http"

	"github.com/anucha-tk/fc_go_scratch/internal/auth"
	"github.com/anucha-tk/fc_go_scratch/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// middleware auth
// check apikey and return user exist
func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApIKey(r.Header)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Error get apikey %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKEY(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 400, fmt.Sprintf("Couldn't get user %v", err))
			return
		}
		handler(w, r, user)
	}
}
