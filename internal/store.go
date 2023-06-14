package internal

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// Storage is a struct that contains the sqlx.DB
type Storage struct {
	DB *sqlx.DB
}

// NewStorage generates a new Storage struct with functioning sqlx.DB.
func NewStorage() (*Storage, error) {
	db, err := newSQLXFromEnv(Config)
	if err != nil {
		return nil, err
	}

	return &Storage{
		DB: db,
	}, nil
}

// Constants used to create a Postgres connection.
const (
	DatasourceNameFormat = "host=%s port=%d dbname=%s user=%s password=%s search_path=%s sslmode=%s"
	PQDriver             = "postgres"
)

func newSQLXFromEnv(cfg Configuration) (*sqlx.DB, error) {
	db, err := sqlx.Open(PQDriver, fmt.Sprintf(DatasourceNameFormat, cfg.PGHost, cfg.PGPort, cfg.PGDatabase, cfg.PGUser, cfg.PGPassword, cfg.PGSchema, cfg.PGSSLMode))
	if err != nil {
		fmt.Println("couldn't open a db connection")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
