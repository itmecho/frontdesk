package server

import (
	"fmt"
	"net/http"

	"github.com/itmecho/frontdesk/pkg/authenticator"
	"github.com/itmecho/frontdesk/pkg/store"
	"github.com/sirupsen/logrus"
)

// Server handles requests and interacts with the backend store
type Server struct {
	*http.Server
	auth   authenticator.Authenticator
	logger *logrus.Logger
	store  store.Store
}

// NewServer returns a new server populated with the given store
func NewServer(port int, logger *logrus.Logger, s store.Store, auth authenticator.Authenticator) Server {
	httpSrv := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	srv := Server{
		Server: httpSrv,
		auth:   auth,
		logger: logger,
		store:  s,
	}

	srv.routes()

	return srv
}
