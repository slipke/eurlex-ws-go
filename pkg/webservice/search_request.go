package webservice

import (
	"encoding/xml"

	"github.com/slipke/eurlex-ws-go/pkg/request"
)

type SearchRequest struct {
	Username       string
	Password       string
	Query          *ExpertQuery
	Page           int64
	PageSize       int64
	SearchLanguage string
}

func NewSearchRequest(query string) *SearchRequest {
	q := NewExpertQuery(query)
	return &SearchRequest{
		Query:          q,
		Page:           1,
		PageSize:       10,
		SearchLanguage: "en",
	}
}

func NewSearchRequestWithConfig(cfg *Config, query string) *SearchRequest {
	r := NewSearchRequest(query)

	r.Username = cfg.Username
	r.Password = cfg.Password
	r.PageSize = cfg.PageSize
	r.SearchLanguage = cfg.SearchLanguage

	return r
}

func (s *SearchRequest) ToXML() ([]byte, error) {
	e := request.NewEnvelope()
	e.Header.RootElement = request.NewSecurity(s.Username, s.Password)
	e.Body.RootElement = request.NewSearchRequest(s.Query.String(), s.Page, s.PageSize, s.SearchLanguage)

	xmlBytes, err := xml.Marshal(e)
	if err != nil {
		return nil, err
	}
	return xmlBytes, nil
}
