package server

import (
	"net/http"

	"github.com/medhir/musicbrainz/server/mbclient"
)

type Server struct {
	Router   *http.ServeMux
	MBClient *mbclient.MBClient
}
