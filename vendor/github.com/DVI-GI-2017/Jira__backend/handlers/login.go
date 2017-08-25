package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/mux"
	"github.com/DVI-GI-2017/Jira__backend/pool"
)

// Registers user
// Post body - user credentials in format: {"email": "...", "password": "..."}
// Returns credentials if OK
func RegisterUser(w http.ResponseWriter, req *http.Request) {
	var credentials models.User

	params := mux.Params(req)

	if err := json.Unmarshal(params.Body, &credentials); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := credentials.Validate(); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	credentials.Encrypt()

	exists, err := pool.Dispatch(pool.UserExists, credentials)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if exists.(bool) {
		JsonErrorResponse(w, fmt.Errorf("User with email: %s already exists!", credentials.Email),
			http.StatusConflict)
		return
	}

	user, err := pool.Dispatch(pool.UserCreate, credentials)
	if err != nil {
		JsonErrorResponse(w, fmt.Errorf("can not create account: %v", err), http.StatusBadGateway)
		return
	}

	token, err := auth.NewToken()
	if err != nil {
		JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	JsonResponse(w, struct {
		models.User
		auth.Token
	}{user.(models.User), token})
}

// Authorizes user in system.
// Post body - credentials in format: {"email": "...", "password": "..."}
// Returns token for authentication.
func Login(w http.ResponseWriter, req *http.Request) {
	var credentials models.User

	params := mux.Params(req)

	if err := json.Unmarshal(params.Body, &credentials); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := credentials.Validate(); err != nil && err != models.ErrEmptyName {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	credentials.Encrypt()

	valid, err := pool.Dispatch(pool.UserAuthorize, credentials)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if !valid.(bool) {
		JsonErrorResponse(w, fmt.Errorf("can not find user with: %s", credentials.Email), http.StatusNotFound)
		return
	}

	user, err := pool.Dispatch(pool.UserFindByEmail, credentials.Email)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	token, err := auth.NewToken()
	if err != nil {
		JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	JsonResponse(w, struct {
		models.User
		auth.Token
	}{user.(models.User), token})
}
