package store

import (
	"errors"

	"github.com/itmecho/frontdesk/pkg/types"
	"github.com/sirupsen/logrus"
)

// Store is an interface for interacting with a backend user store
type Store interface {
	Migrate() error

	GetAll() ([]types.User, error)
	GetByID(id string) (*types.User, error)
	Create(*types.User) error
	Update(*types.User) error
	Delete(*types.User) error
}

// NewStore returns a new store based on the provided database type
func NewStore(dbType, dsn string, logger *logrus.Logger) (Store, error) {
	switch dbType {
	case "postgres":
		return newPostgresStore(dsn, logger)
	case "cockroach":
		return newPostgresStore(dsn, logger)
	default:
		return nil, errors.New("unsupported database type: " + dbType)
	}
}
