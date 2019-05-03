package mbclient

import "encoding/xml"

const ReleaseEntity = "release"

type ReleaseMetadata struct {
	XMLName     xml.Name    `xml:"metadata" json:"-"`
	ReleaseList releaseList `xml:"release-list"`
}

type releaseList struct {
	XMLName  xml.Name  `xml:"release-list" json:"-"`
	Count    string    `xml:"count,attr" json:"count"`
	Releases []release `xml:"release" json:"releases"`
}

type release struct {
	XMLName        xml.Name     `xml:"release" json:"-"`
	ID             string       `xml:"id,attr" json:"id"`
	Title          string       `xml:"title" json:"title"`
	Status         string       `xml:"status" json:"status"`
	Disambiguation string       `xml:"disambiguation" json:"disambiguation"`
	Credit         artistCredit `xml:"artist-credit" json:"-"`
}

type artistCredit struct {
	XMLName    xml.Name   `xml:"artist-credit" json:"-"`
	NameCredit nameCredit `xml:"name-credit"`
}

type nameCredit struct {
	XMLName xml.Name `xml:"name-credit" json:"-"`
	Artist  artist   `xml:"artist" json:"artist"`
}

func (c *MBClient) GetReleasesByArtist(id string, typeFilters []string) (*ReleaseMetadata, error) {
	metadata := &ReleaseMetadata{}
	q := c.CreateQuery()
	q.Set("artist", id)
	if len(typeFilters) > 0 {
		types := c.getTypeString(typeFilters)
		q.Set("type", types)
	}
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

func (c *MBClient) GetReleasesByArtistAndTitle(id, title string, typeFilters []string) (*ReleaseMetadata, error) {
	idQuery := "arid:" + id
	finalQuery := title + " AND " + idQuery
	if len(typeFilters) > 0 {
		types := c.getTypeString(typeFilters)
		finalQuery += " AND " + types
	}
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
