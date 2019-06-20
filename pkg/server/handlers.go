package server

import (
	"encoding/json"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/itmecho/frontdesk/pkg/types"
)

func (srv *Server) getUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		out := json.NewEncoder(w)

		users, err := srv.store.GetAll()
		if err != nil {
			srv.logger.Error(err)
		}

		out.Encode(users)
	}
}

func (srv *Server) createUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestDecoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		type requestObject struct {
			Name                 string `json:"name"`
			Email                string `json:"email"`
			Password             string `json:"password"`
			PasswordConfirmation string `json:"password_confirmation"`
		}

		var newUserRequest requestObject

		if err := requestDecoder.Decode(&newUserRequest); err != nil {
			srv.logger.Error("failed to decode request: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newUserRequest.Password != newUserRequest.PasswordConfirmation {
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
