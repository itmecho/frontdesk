package server

import "github.com/go-chi/chi"

func (srv *Server) routes() {
	router := chi.NewRouter()
	router.Route("/api/", func(r chi.Router) {
		r.Get("/users", srv.getUsersHandler())
	})

	srv.Handler = router
}
