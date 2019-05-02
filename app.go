package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/medhir/musicbrainz/server/mbclient"
)

const BaseURL = "https://musicbrainz.org/ws/2/"
const UserAgent = "Medhir's Musicbrainz Client App / v0.1 / Contact: mail AT medhir.com"

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10}
	parsedURL, _ := url.Parse(BaseURL)
	client := &mbclient.MBClient{
		BaseURL:    parsedURL,
		UserAgent:  UserAgent,
		HTTPClient: httpClient}
	id, err := client.GetFirstArtistID("Drake")
	if err != nil {
		fmt.Println(err)
	}
	releases, err := client.GetReleasesByArtistAndTitle(id, "more")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Drake Titles")
	for _, release := range releases.ReleaseList.Releases {
		fmt.Println("-> " + release.Title)
	}
}
