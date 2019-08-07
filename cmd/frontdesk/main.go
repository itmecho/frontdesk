package main

import (
	"github.com/itmecho/frontdesk/pkg/authenticator"
	"github.com/itmecho/frontdesk/pkg/server"
	"github.com/itmecho/frontdesk/pkg/store"
	"github.com/itmecho/frontdesk/pkg/store/postgres"
)

func main() {
	var db store.Store
	var err error

	switch databaseType {
	case "postgres":
		db, err = postgres.NewStore(databaseDSN, logger)
	default:
		logger.Fatal("Unknown database type: ", databaseType)
	}

	if err != nil {
		logger.Fatal("Failed to create the database connection: ", err)
	}

	if err = db.Migrate(); err != nil {
		logger.Fatal("Failed to migrate the database: ", err)
	}

	auth, err := authenticator.New([]byte(authSecret))
	if err != nil {
		logger.Fatal("Failed to create authenticator: ", err)
	}

	srv := server.NewServer(port, logger, db, auth)

	logger.Infof("ready to handle requests on port %d", port)
	if err = srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
