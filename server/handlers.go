package server

import (
	"encoding/json"
	"net/http"
)

func (s *server) handleSearch() http.HandlerFunc {
	type searchRequestBody struct {
		Artist  string   `json:"artist"`
		Title   string   `json:"title"`
		Filters []string `json:"filters"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		searchRequest := &searchRequestBody{}
		json.NewDecoder(r.Body).Decode(searchRequest)
		artist, err := s.MBClient.GetFirstArtistID(searchRequest.Artist)
		if err != nil {
			http.Error(w, "Could not get first artist => "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if searchRequest.Title != "" {
			response, err := s.MBClient.GetReleasesByArtistAndTitle(artist, searchRequest.Title, searchRequest.Filters)
			if err != nil {
				http.Error(w, "Could not perform search query "+err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(response)
		} else {
			response, err := s.MBClient.GetReleasesByArtist(artist, searchRequest.Filters)
			if err != nil {
				http.Error(w, "Could not perform search query "+err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}

func (s *server) handleGetAlbumInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
