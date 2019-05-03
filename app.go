package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/medhir/musicbrainz/server"

	"github.com/medhir/musicbrainz/server/mbclient"
)

// BaseURL is the API Endpoint for the Musicbrainz client
const BaseURL = "https://musicbrainz.org/ws/2/"

// UserAgent provides a description of the application to be sent with Musicbrainz API requests
const UserAgent = "Medhir's Musicbrainz Client App / v0.1 / Contact: mail AT medhir.com"

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10}
	parsedURL, _ := url.Parse(BaseURL)
	client := &mbclient.MBClient{
		BaseURL:    parsedURL,
		UserAgent:  UserAgent,
		HTTPClient: httpClient}
	mux := http.NewServeMux()
	server := server.NewServer(mux, client)
	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", server.Router)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
