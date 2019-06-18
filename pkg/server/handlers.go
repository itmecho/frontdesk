package server

import (
	"encoding/json"
	"net/http"
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
