package types

import "time"

type URL struct {
	ID         int       `json:"id"`
	CustomSlug string    `json:"custom_slug"`
	LongURL    string    `json:"long_url"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiresAt  time.Time `json:"expires_at"`
}

type CreateURLRequest struct {
	CustomSlug string `json:"custom_slug"`
	LongURL    string `json:"long_url"`
}

func NewURL(customSlug, longURL string) (*URL, error) {
	return &URL{
		CustomSlug: customSlug,
		LongURL:    longURL,
		CreatedAt:  time.Now(),
		ExpiresAt:  time.Now().Add(time.Hour * 24 * 7),
	}, nil
}
