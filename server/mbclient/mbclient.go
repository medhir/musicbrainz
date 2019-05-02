package mbclient

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

// MBClient describes parameters of a Musicbrainz client
type MBClient struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	UserAgent  string
}

// NewRequest constructs a request to be sent to the Musicbrainz API
func (c *MBClient) NewRequest(entity, query string) (*http.Request, error) {
	rel := &url.URL{Path: entity}
	u := c.BaseURL.ResolveReference(rel)
	q := u.Query()
	q.Set("query", query)
	u.RawQuery = q.Encode()
	// var buf io.ReadWriter
	// if body != nil {
	// 	buf = new(bytes.Buffer)
	// 	err := json.NewEncoder(buf).Encode(body)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	fmt.Println(u.String())
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

// Do executes a request and decodes the response into v
func (c *MBClient) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = xml.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
