package store

import (
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

// MigrationLogger implements the migrate.Logger interface
type MigrationLogger struct {
	*logrus.Logger
}

// Verbose checks if verbose logging should be used
func (l *MigrationLogger) Verbose() bool {
	return l.IsLevelEnabled(logrus.DebugLevel)
}

// WrapMigrationLogger returns a logger that can be used during migrations
func WrapMigrationLogger(l *logrus.Logger) *MigrationLogger {
	return &MigrationLogger{l}
}
