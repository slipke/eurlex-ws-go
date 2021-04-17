package response

import (
	"encoding/xml"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  *Header  `xml:"Header"`
	Body    *Body    `xml:"Body"`
}

func NewEnvelopeFromXML(xmlStr string) (*Envelope, error) {
	e := &Envelope{}
	if err := xml.Unmarshal([]byte(xmlStr), e); err != nil {
		return nil, err
	}
	return e, nil
}

type Header struct {
	XMLName       xml.Name       `xml:"Header"`
	NotUnderstood *NotUnderstood `xml:"NotUnderstood"`
}

type NotUnderstood struct {
	XMLName xml.Name `xml:"NotUnderstood"`
	Qname   string   `xml:"qname,attr"`
}

type Body struct {
	XMLName       xml.Name       `xml:"Body"`
	Fault         *Fault         `xml:"Fault"`
	SearchResults *SearchResults `xml:"searchResults"`
}

func NewBody() *Body {
	return &Body{}
}

type Fault struct {
	XMLName xml.Name `xml:"Fault"`
	Code    string   `xml:"Code>Value"`
	Reason  string   `xml:"Reason>Text"`
}

type SearchResults struct {
	XMLName   xml.Name `xml:"searchResults"`
	NumHits   int64    `xml:"numhits"`
	TotalHits int64    `xml:"totalhits"`
	Page      int64    `xml:"page"`
	Language  string   `xml:"language"`
	Result    *Result  `xml:"result"`
}

func NewSearchResults() *SearchResults {
	return &SearchResults{}
}

type Result struct {
	// @TODO
}
