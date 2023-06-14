package internal

import (
	"context"
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

// SaveJournal stores new journal in DB.
func (s *Storage) SaveJournal(ctx context.Context, j Journal) (*Journal, error) {
	var newJournal Journal
	err := s.DB.QueryRowxContext(ctx, insertJournalQuery,
		j.Name,
	).StructScan(&newJournal)
	if err != nil {
		fmt.Println("error saving journal in DB")
		return nil, err
	}

	return &newJournal, nil
}

// RetrieveAllJournals gets all journals from the DB
func (s *Storage) RetrieveAllJournals(ctx context.Context) (*[]Journal, error) {
	rows, err := s.DB.QueryxContext(ctx,
		retrieveAllJournalsQuery,
	)
	if err != nil {
		fmt.Println("error retrieving journals from db")
		return nil, err
	}
	defer rows.Close()

	journalList := make([]Journal, 0)
	j := Journal{}
	for rows.Next() {
		rows.StructScan(&j)
		journalList = append(journalList, j)
	}

	return &journalList, nil
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
