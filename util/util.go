package util

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/ybencab/url-shortener/types"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func ReadJSON(r *http.Request, data any) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func ScanRowIntoURL(row *sql.Rows) (*types.URL, error) {
	url := new(types.URL)
	err := row.Scan(
		&url.ID,
		&url.CustomSlug,
		&url.LongURL,
		&url.CreatedAt,
		&url.ExpiresAt,
	)
	return url, err
}

func IsValidURL(urlString string) bool {
	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	resp, err := http.Head(urlString)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode >= 200 && resp.StatusCode < 300
}
