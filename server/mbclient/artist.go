package mbclient

import "encoding/xml"

const ArtistEntity = "artist"

type ArtistMetadata struct {
	XMLName    xml.Name   `xml:"metadata" json:"-"`
	ArtistList artistList `xml:"artist-list"`
}

type artistList struct {
	XMLName xml.Name `xml:"artist-list" json:"-"`
	Count   string   `xml:"count,attr" json:"count"`
	Artists []artist `xml:"artist" json:"artists"`
}

type artist struct {
	XMLName xml.Name `xml:"artist"`
	ID      string   `xml:"id,attr" json:"id"`
	Name    string   `xml:"name" json:"name"`
}

// GetArtistMetadata does a search for artists by a string query, returns a response encoded as an ArtistMetadata struct
func (c *MBClient) GetArtistMetadata(artistQuery string) (*ArtistMetadata, error) {
	metadata := &ArtistMetadata{}
	q := c.CreateQuery()
	q.Set("query", artistQuery)
	req, err := c.NewRequest(ArtistEntity, q)
	if err != nil {
		return nil, err
	}
	_, err = c.Do(req, metadata)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

// GetFirstArtistID returns the first artist match's Musicbrainz ID value
func (c *MBClient) GetFirstArtistID(artistQuery string) (string, error) {
	metadata := &ArtistMetadata{}
	q := c.CreateQuery()
	q.Set("query", artistQuery)
	req, err := c.NewRequest(ArtistEntity, q)
	if err != nil {
		return "", err
	}
	_, err = c.Do(req, metadata)
	if err != nil {
		return "", err
	}
	return metadata.ArtistList.Artists[0].ID, nil
}
