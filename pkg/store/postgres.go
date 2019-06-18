package store

import (
	"github.com/itmecho/frontdesk/pkg/types"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// PostgresStore is an implementation of Store to interact with a postgres compatible backend
type PostgresStore struct {
	db     *sqlx.DB
	logger *logrus.Logger
}

func newPostgresStore(dsn string, logger *logrus.Logger) (Store, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &PostgresStore{
		db:     db,
		logger: logger,
	}, err
}

// GetAll loads all users from the store
func (s *PostgresStore) GetAll() ([]types.User, error) {
	users := make([]types.User, 0)

	s.logger.Info("loading all users from the store")

	return users, nil
}

// GetByID loads a user from the store by their ID
func (s *PostgresStore) GetByID(id string) (types.User, error) {
	var user types.User

	s.logger.Info("loading user with id ", id)

	return user, nil
}

// Create inserts a new user into the store
func (s *PostgresStore) Create(user types.User) error {
	s.logger.Info("creating new user")

	return nil
}

// Update saves an existing user into the database with any modifications
func (s *PostgresStore) Update(user types.User) error {
	s.logger.Info("updating user with id ", user.ID)
	return nil
}

// Delete marks a user in the store as deleted
func (s *PostgresStore) Delete(user types.User) error {
	s.logger.Info("deleting user with id ", user.ID)
	return nil
}