package storage

import "github.com/ybencab/url-shortener/types"

type Storer interface {
	GetURLs() ([]*types.URL, error)
	GetURLBySlug(string) (*types.URL, error)
	CreateURL(*types.URL) error
}
