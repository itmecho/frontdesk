package postgres

import (
	"errors"

	"github.com/itmecho/frontdesk/pkg/store"
	"github.com/itmecho/frontdesk/pkg/store/postgres/migrations"
	"github.com/itmecho/frontdesk/pkg/types"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// Store is an implementation of Store to interact with a postgres compatible backend
type Store struct {
	db     *sqlx.DB
	logger *logrus.Logger
}

// NewStore returns a new instance of a postgres store
func NewStore(dsn string, logger *logrus.Logger) (store.Store, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:     db,
		logger: logger,
	}, err
}

// GetAll loads all users from the store
func (s *Store) GetAll() ([]types.User, error) {
	users := make([]types.User, 0)
	s.logger.Info("loading all users from the store")
	err := s.db.Select(&users, "SELECT * FROM users")

	return users, err
}

// GetByID loads a user from the store by their ID
func (s *Store) GetByID(id string) (*types.User, error) {
	s.logger.Info("loading user with id ", id)

	var user types.User

	if err := s.db.Get(&user, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetByEmail loads a user from the store by their email
func (s *Store) GetByEmail(email string) (*types.User, error) {
	s.logger.Info("loading user with email ", email)

	var user types.User

	if err := s.db.Get(&user, "SELECT * FROM users WHERE email = $1", email); err != nil {
		return nil, err
	}

	return &user, nil
}

// Create inserts a new user into the store
func (s *Store) Create(user *types.User) error {
	s.logger.Info("creating new user")

	_, err := s.db.Exec("INSERT INTO users (id, name, email, password) values($1, $2, $3, $4)",
		user.ID,
		user.Name,
		user.Email,
		user.Password,
	)
	if err != nil {
		return s.wrapError(err)
	}

	return nil
}

// Update saves an existing user into the database with any modifications
func (s *Store) Update(user *types.User) error {
	s.logger.Info("updating user with id ", user.ID)
	return nil
}

// Delete marks a user in the store as deleted
func (s *Store) Delete(user *types.User) error {
	s.logger.Info("deleting user with id ", user.ID)
	return nil
}

func (s *Store) wrapError(err error) error {
	if pqErr, ok := err.(*pq.Error); ok {
		s.logger.Info(pqErr.Code.Name())
		switch pqErr.Code.Name() {
		case "unique_violation":
			switch pqErr.Constraint {
			case "users_email_key":
				return errors.New("A user with that email address already exists")
			}
		}
	}

	return err
}

// Migrate brings the database up to the latest version
func (s *Store) Migrate() error {
	migrateDB, err := migratePostgres.WithInstance(s.db.DB, &migratePostgres.Config{})
	if err != nil {
		return err
	}

	source := bindata.Resource(migrations.AssetNames(), migrations.Asset)
	data, err := bindata.WithInstance(source)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("go-bindata", data, "postgres", migrateDB)
	if err != nil {
		return err
	}
	m.Log = store.WrapMigrationLogger(s.logger)
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			s.logger.Info("Database is up to date")
		} else {
			return err
		}
	}

	return err
}
