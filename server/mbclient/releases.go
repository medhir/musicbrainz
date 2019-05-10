package mbclient

import (
	"encoding/xml"
)

const ReleaseEntity = "release"

type ReleasesMetadata struct {
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

func (c *MBClient) GetReleasesByArtist(id string, typeFilters []string, offset string) (*ReleasesMetadata, error) {
	finalQuery := c.createReleasesQuery(id, "", typeFilters)
	metadata, err := c.executeReleasesQuery(finalQuery, offset)
	return metadata, err
}

func (c *MBClient) GetReleasesByArtistAndTitle(id, title string, typeFilters []string, offset string) (*ReleasesMetadata, error) {
	finalQuery := c.createReleasesQuery(id, title, typeFilters)
	metadata, err := c.executeReleasesQuery(finalQuery, offset)
	return metadata, err
}

func (c *MBClient) createReleasesQuery(id, title string, typeFilters []string) string {
	idQuery := "arid:" + id
	var finalQuery string
	if title == "" {
		finalQuery = idQuery + " AND country:US"
	} else {
		finalQuery = title + " AND " + idQuery + " AND country:US"
	}
	if len(typeFilters) > 0 {
		types := c.getTypeString(typeFilters)
		finalQuery += " AND " + types
	}
	return finalQuery
}

func (c *MBClient) executeReleasesQuery(query string, offset string) (*ReleasesMetadata, error) {
	q := c.CreateQuery()
	q.Set("query", query)
	q.Set("offset", string(offset))
	metadata := &ReleasesMetadata{}
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
