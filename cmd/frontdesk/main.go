package main

import (
	"github.com/itmecho/frontdesk/pkg/server"
	"github.com/itmecho/frontdesk/pkg/store"
)

func main() {
	db, err := store.NewStore(databaseType, databaseDSN, logger)
	if err != nil {
		logger.Fatal(err)
	}

	if err := db.Migrate(); err != nil {
		logger.Fatal("Failed to migrate store: ", err)
	}

	srv := server.NewServer(port, logger, db)

	logger.Infof("ready to handle requests on port %d", port)
	if err = srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
