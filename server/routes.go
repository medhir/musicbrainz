package server

func (s *server) routes() {
	s.Router.HandleFunc("/api/search", s.handleSearch())
	s.Router.HandleFunc("/api/album", s.handleGetAlbumInfo())
}
