package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var (
	databaseType string
	databaseDSN  string

	port int

	logger *logrus.Logger
)

func init() {
	pflag.StringVar(&databaseType, "database-type", "postgres", "Database type to use")
	pflag.StringVar(&databaseDSN, "database-dsn", "host=localhost port=26257 dbname=frontdesk user=root sslmode=disable", "Database connection string")
	pflag.IntVar(&port, "port", 9000, "Port for the server to listen on")

	pflag.Parse()

	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
}
