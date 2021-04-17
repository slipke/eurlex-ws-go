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
		Query: q,
	}
}

func NewSearchRequestWithConfig(cfg *Config, query string) *SearchRequest {
	r := NewSearchRequest(query)

	r.Username = cfg.Username
	r.Password = cfg.Password

	return r
}

func (s *SearchRequest) ToXML() ([]byte, error) {
	e := request.NewEnvelope()

	e.Header.Security.UsernameToken.Username = s.Username
	e.Header.Security.UsernameToken.Password.Password = s.Password

	sr := request.NewSearchRequest(s.Query.String(), s.Page, s.PageSize, s.SearchLanguage)
	e.Body.RootElement = sr

	xmlBytes, err := xml.Marshal(e)
	if err != nil {
		return nil, err
	}
	return xmlBytes, nil
}
