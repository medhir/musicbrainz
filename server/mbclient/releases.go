package mbclient

import "encoding/xml"

const ReleaseEntity = "release"

type ReleaseMetadata struct {
	XMLName     xml.Name    `xml:"metadata"`
	ReleaseList releaseList `xml:"release-list"`
}

type releaseList struct {
	XMLName  xml.Name  `xml:"release-list"`
	Count    string    `xml:"count,attr"`
	Releases []release `xml:"release"`
}

type release struct {
	XMLName        xml.Name     `xml:"release"`
	ID             string       `xml:"id, attr"`
	Title          string       `xml:"title"`
	Status         string       `xml:"status"`
	Disambiguation string       `xml:"disambiguation"`
	Credit         artistCredit `xml:"artist-credit"`
}

type artistCredit struct {
	XMLName    xml.Name   `xml:"artist-credit"`
	NameCredit nameCredit `xml:"name-credit"`
}

type nameCredit struct {
	XMLName xml.Name `xml:"name-credit"`
	Artist  artist   `xml:"artist"`
}

func (c *MBClient) GetReleasesByArtist(id string) (*ReleaseMetadata, error) {
	metadata := &ReleaseMetadata{}
	q := c.CreateQuery()
	q.Set("artist", id)
	req, err := c.NewRequest(ReleaseEntity, q)
	if err != nil {
		return nil, err
	}
	_, err = c.Do(req, metadata)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

func (c *MBClient) GetReleasesByArtistAndTitle(id, title string) (*ReleaseMetadata, error) {
	idQuery := "arid:" + id
	finalQuery := title + " AND " + idQuery
	q := c.CreateQuery()
	q.Set("query", finalQuery)
	metadata := &ReleaseMetadata{}
	req, err := c.NewRequest(ReleaseEntity, q)
	if err != nil {
		return nil, err
	}
	_, err = c.Do(req, metadata)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}
