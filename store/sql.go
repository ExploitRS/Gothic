package store

import (
	"database/sql"
	"fmt"

	"github.com/exploit-rs/gothic/config/database"
	_ "github.com/jackc/pgx/v5"
)

type DB struct {
	DBInstance *sql.DB
	config     *database.Database
}

func New(cfg *database.Database) *DB {
	db := &DB{
		config: cfg,
	}
	return db
}

func (db *DB) Open() (err error) {
	if db.config.IsDSNEmpty() {
		return fmt.Errorf("dsn required")
	}

	// Connect to the database
	postgre, err := sql.Open("postgres",
		fmt.Sprintf("postgresql://%s:%s@%s:%v",
			db.config.Username(),
			db.config.Password(),
			db.config.Address(),
			db.config.Port(),
		),
	)
	if err != nil {
		return fmt.Errorf("failed to open db with dsn: %s", db.config.DSN())
	}
	db.DBInstance = postgre
	return nil
}
