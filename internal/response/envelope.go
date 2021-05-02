package response

import (
	"encoding/xml"

	model "github.com/slipke/eurlex-model-go"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  *Header  `xml:"Header"`
	Body    *Body    `xml:"Body"`
}

func NewEnvelopeFromXML(xmlStr string) (*Envelope, error) {
	var e *Envelope
	if err := xml.Unmarshal([]byte(xmlStr), &e); err != nil {
		return nil, err
	}
	return e, nil
}

type Header struct {
	NotUnderstood *NotUnderstood `xml:"NotUnderstood"`
}

type NotUnderstood struct {
	Qname string `xml:"qname,attr"`
}

type Body struct {
	Fault         *Fault         `xml:"Fault"`
	SearchResults *SearchResults `xml:"searchResults"`
}

type Fault struct {
	Code   string `xml:"Code>Value"`
	Reason string `xml:"Reason>Text"`
}

type SearchResults struct {
	NumHits   int64     `xml:"numhits"`
	TotalHits int64     `xml:"totalhits"`
	Page      int64     `xml:"page"`
	Language  string    `xml:"language"`
	Result    []*Result `xml:"result"`
}

type Result struct {
	Reference string   `xml:"reference"`
	Rank      string   `xml:"rank"`
	Content   *Content `xml:"content"`
}

type Content struct {
	Notice *model.Notice `xml:"NOTICE"`
}
