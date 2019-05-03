package mbclient

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

// MBClient describes parameters of a Musicbrainz client
type MBClient struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	UserAgent  string
}

func (c *MBClient) CreateQuery() *url.Values {
	u := &url.URL{}
	q := u.Query()
	return &q
}

// NewRequest constructs a request to be sent to the Musicbrainz API
func (c *MBClient) NewRequest(entity string, q *url.Values) (*http.Request, error) {
	rel := &url.URL{Path: entity}
	u := c.BaseURL.ResolveReference(rel)
	u.RawQuery = q.Encode()
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
