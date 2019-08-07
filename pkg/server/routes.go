package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (srv *Server) routes() {
	router := chi.NewRouter()

	router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  srv.logger,
		NoColor: true,
	}))

	router.Route("/api/", func(r chi.Router) {
		r.Get("/users", srv.getUsersHandler())
		r.Post("/users", srv.createUserHandler())
		r.Post("/authenticate", srv.authenticateHandler())
	})

	srv.Handler = router
}
