package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ybencab/url-shortener/types"
	"github.com/ybencab/url-shortener/util"
)

func (s *Server) handleGetURLs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "method not allowed"})
		return
	}

	urls, err := s.store.GetURLs()
	if err != nil {
		util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	util.WriteJSON(w, http.StatusOK, urls)
}

func (s *Server) handleCreateURLs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "method not allowed"})
		return
	}

	urlReq := new(types.CreateURLRequest)
	if err := util.ReadJSON(r, urlReq); err != nil {
		util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if len(urlReq.CustomSlug) < 5 || len(urlReq.CustomSlug) > 20 {
		util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid slug"})
		return
	}
	
	if _, err := s.store.GetURLBySlug(urlReq.CustomSlug); err == nil {
		util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "slug already exists"})
		return
	}

	if !util.IsValidURL(urlReq.LongURL) {
		util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid url"})
		return
	}

	newUrl, err := types.NewURL(urlReq.CustomSlug, urlReq.LongURL)
	if err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "server error"})
		return
	}

	if err := s.store.CreateURL(newUrl); err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "database error"})
		return
	}

	util.WriteJSON(w, http.StatusOK, map[string]string{"success": "url created"})
}

func (s *Server) handleGetURLBySlug(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "method not allowed"})
		return
	}

	slug := mux.Vars(r)["slug"]
	url, err := s.store.GetURLBySlug(slug)
	if err != nil {
		util.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "url does not exist"})
		return
	}

	util.WriteJSON(w, http.StatusOK, url)
}
