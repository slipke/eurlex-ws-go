package response

import (
	"encoding/xml"
)

type Envelope struct {
	XMLName       xml.Name `xml:"Envelope"`
	XMLNamespaceS string   `xml:"S,attr"`
	Body          *Body    `xml:"Body"`
}

func NewEnvelopeFromXML(xmlStr string) (*Envelope, error) {
	e := &Envelope{}

	if err := xml.Unmarshal([]byte(xmlStr), e); err != nil {
		return nil, err
	}

	return e, nil
}

type Body struct {
	XMLName       xml.Name      `xml:"Body"`
	SearchResults *SearchResult `xml:"searchResults"`
}

func NewBody() *Body {
	return &Body{}
}

type SearchResult struct {
	XMLName         xml.Name `xml:"searchResults"`
	SchemaLocation  string   `xml:"xsi:schemaLocation,attr"`
	XMLNamespace    string   `xml:"xmlns,attr"`
	XMLNamespaceXSI string   `xml:"xmlns:xsi,attr"`
	NumHits         int64    `xml:"numhits"`
	TotalHits       int64    `xml:"totalhits"`
	Page            int64    `xml:"page"`
	Language        string   `xml:"language"`
	Result          *Result  `xml:"result"`
}

func NewSearchResult() *SearchResult {
	return &SearchResult{}
}

type Result struct {
	// @TODO
}
