package eurlex

import (
	"encoding/xml"

	"github.com/slipke/eurlex-ws-go/internal/request"
)

type SearchRequest struct {
	username       string
	password       string
	Query          *ExpertQuery
	Page           int64
	PageSize       int64
	SearchLanguage string
}

func NewSearchRequest() *SearchRequest {
	return &SearchRequest{
		Page:           1,
		PageSize:       10,
		SearchLanguage: "en",
	}
}

func NewSearchRequestFromString(query string) *SearchRequest {
	q := NewExpertQueryFromString(query)
	sr := NewSearchRequest()
	sr.Query = q
	return sr
}

func (s *SearchRequest) ToXML() ([]byte, error) {
	e := request.NewEnvelope()
	e.Header.RootElement = request.NewSecurity(s.username, s.password)

	if s.Query != nil {
		e.Body.RootElement = request.NewSearchRequest(s.Query.String(), s.Page, s.PageSize, s.SearchLanguage)
	}

	xmlBytes, err := xml.Marshal(e)
	if err != nil {
		return nil, err
	}
	return xmlBytes, nil
}
