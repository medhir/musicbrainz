package server

import (
	"net/http"

	"github.com/medhir/musicbrainz/server/mbclient"
)

type server struct {
	Router   *http.ServeMux
	MBClient *mbclient.MBClient
}

// NewServer initializes the server
func NewServer(router *http.ServeMux, client *mbclient.MBClient) *server {
	server := server{
		Router:   router,
		MBClient: client}
	server.routes()
	return &server
}
