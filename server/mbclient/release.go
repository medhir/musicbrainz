package mbclient

import "encoding/xml"

type ReleaseMetadata struct {
	XMLName xml.Name          `xml:"metadata" json:"-"`
	Release releaseWithTracks `xml:"release"`
}

type releaseWithTracks struct {
	XMLName        xml.Name     `xml:"release" json:"-"`
	ID             string       `xml:"id,attr" json:"id"`
	Title          string       `xml:"title" json:"title"`
	Status         string       `xml:"status" json:"status"`
	Disambiguation string       `xml:"disambiguation" json:"disambiguation"`
	Credit         artistCredit `xml:"artist-credit" json:"-"`
	MediumList     mediumList   `xml:"medium-list"`
}

type mediumList struct {
	XMLName xml.Name `xml:"medium-list" json:"-"`
	Mediums []medium `xml:"medium"`
}

type medium struct {
	XMLName   xml.Name  `xml:"medium" json:"-"`
	Position  int       `xml:"position"`
	Format    string    `xml:"format"`
	TrackList trackList `xml:"track-list"`
}

type trackList struct {
	XMLName xml.Name `xml:"track-list" json:"-"`
	Count   int      `xml:"count,attr"`
	Tracks  []track  `xml:"track"`
}

type track struct {
	XMLName   xml.Name  `xml:"track" json:"-"`
	ID        string    `xml:"id,attr"`
	Position  int       `xml:"position"`
	Number    int       `xml:"number"`
	Recording recording `xml:"recording"`
}

type recording struct {
	XMLName        xml.Name `xml:"recording" json:"-"`
	ID             string   `xml:"id,attr"`
	Title          string   `xml:"title"`
	Length         string   `xml:"length"`
	Disambiguation string   `xml:"disambiguation"`
}

// GetReleaseInfo gets the release's information for a release ID. This information includes track lists
func (c *MBClient) GetReleaseInfo(id string) (*ReleaseMetadata, error) {
	endpoint := ReleaseEntity + "/" + id
	q := c.CreateQuery()
	q.Set("inc", "recordings")
	req, err := c.NewRequest(endpoint, q)
	if err != nil {
		return nil, err
	}
	metadata := &ReleaseMetadata{}
	_, err = c.Do(req, metadata)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}
