package storage

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
	"github.com/ybencab/url-shortener/types"
	"github.com/ybencab/url-shortener/util"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error: ", err.Error())
		return nil, nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error: ", err.Error())
		return nil, nil
	}

	return &PostgresStorage{db}, nil
}

func (s *PostgresStorage) GetURLs() ([]*types.URL, error) {
	rows, err := s.db.Query("select * from urls limit 10")
	if err != nil {
		return nil, err
	}

	urls := []*types.URL{}
	for rows.Next() {
		url, err := util.ScanRowIntoURL(rows)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func (s *PostgresStorage) GetURLBySlug(slug string) (*types.URL, error) {
	rows, err := s.db.Query("select * from urls where custom_slug = $1", slug)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return util.ScanRowIntoURL(rows)
	}

	return nil, errors.New("url not found")
}

func (s *PostgresStorage) CreateURL(u *types.URL) error {
	query := `insert into urls
		(created_at, expires_at, long_url, custom_slug)
		values ($1, $2, $3, $4)`

	_, err := s.db.Query(
		query,
		u.CreatedAt,
		u.ExpiresAt,
		u.LongURL,
		u.CustomSlug,
	)
	if err != nil {
		return err
	}

	return nil
}
