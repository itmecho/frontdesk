package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/itmecho/frontdesk/pkg/types"
)

func (srv *Server) handleUsersGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		out := json.NewEncoder(w)

		users, err := srv.store.GetAll()
		if err != nil {
			srv.logger.Error(err)
		}

		out.Encode(users)
	}
}

func (srv *Server) handleUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestDecoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		type requestObject struct {
			Name                 string          `json:"name"`
			Email                string          `json:"email"`
			Password             json.RawMessage `json:"password"`
			PasswordConfirmation json.RawMessage `json:"password_confirmation"`
		}

		var newUserRequest requestObject

		if err := requestDecoder.Decode(&newUserRequest); err != nil {
			srv.logger.Error("failed to decode request: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if bytes.Compare(newUserRequest.Password, newUserRequest.PasswordConfirmation) != 0 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"error":"password and password_confirmation do not match"}`))
			return
		}

		newUUID := uuid.NewV4()

		newUser := &types.User{
			ID:    newUUID.String(),
			Name:  newUserRequest.Name,
			Email: newUserRequest.Email,
		}
		newUser.SetPassword(newUserRequest.Password)

		if err := srv.store.Create(newUser); err != nil {

			srv.logger.Error("failed to create new user: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		newUser, err := srv.store.GetByID(newUUID.String())
		if err != nil {
			srv.logger.Error("failed to load newly created user: ", err)
			// TODO: Should return a better error as it has actually created the user. Maybe wrap in a transaction and roll it back
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		responseEncoder := json.NewEncoder(w)
		// TODO figure out how to handle the error from the following
		responseEncoder.Encode(newUser)

	}
}

func (srv *Server) handleAuthenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestDecoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		type requestObject struct {
			Email    string          `json:"email"`
			Password json.RawMessage `json:"password"`
		}

		authenticateRequest := requestObject{}

		if err := requestDecoder.Decode(&authenticateRequest); err != nil {
			srv.logger.Error("failed to decode request: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := srv.store.GetByEmail(authenticateRequest.Email)
		if err != nil {
			srv.logger.Error("failed to find user by email: ", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"A user with that email was not found"}`))
			return
		}

		if !user.CheckPassword(authenticateRequest.Password) {
			srv.logger.Error("password check failed")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"A user with that email and password was not found"}`))
			return
		}

		token, err := srv.auth.GenerateToken(user)
		if err != nil {
			srv.logger.Error("failed to generate JWT: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"Failed to create authentication token"}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("{\"token\":\"%s\"}", token)))
	}
}
