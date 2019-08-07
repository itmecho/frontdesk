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
		r.Get("/users", srv.handleUsersGet())
		r.Post("/users", srv.handleUserCreate())
		r.Post("/authenticate", srv.handleAuthenticate())
	})

	srv.Handler = router
}
