package store

// Migrate brings the postgres store up to date
func (pgStore PostgresStore) Migrate() error {
	pgStore.logger.Info("beginning store migration")

	for _, mig := range postgresMigrations {
		pgStore.logger.Info("executing migration: ", mig.Name)
		if _, err := pgStore.db.Exec(mig.Up); err != nil {
			return err
		}
	}

	return nil
}

var postgresMigrations = []migration{
	{
		Name: "create_user_table",
		Up: `CREATE TABLE IF NOT EXISTS users (
			id uuid,
			name VARCHAR NOT NULL,
			email VARCHAR UNIQUE NOT NULL,
			password VARCHAR NOT NULL,
			created_at TIMESTAMP DEFAULT now(),
			updated_at TIMESTAMP DEFAULT now(),
			deleted_at TIMESTAMP,
			PRIMARY KEY (id)
		);`,
		Down: `DROP TABLE users;`,
	},
}
