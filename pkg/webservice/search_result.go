package webservice

import (
	"fmt"

	"github.com/slipke/eurlex-ws-go/pkg/response"
)

type SearchResult struct {
	NumHits   int64
	TotalHits int64
	Page      int64
	Language  string
	Result    []*response.Result
}

func NewSearchResult() *SearchResult {
	return &SearchResult{}
}

func NewSearchResultFromXML(xml string) (*SearchResult, error) {
	sr := NewSearchResult()
	e, err := response.NewEnvelopeFromXML(xml)
	if err != nil {
		return nil, fmt.Errorf("failed to create envelope from XML: %s", err)
	}

	sr.NumHits = e.Body.SearchResults.NumHits
	sr.TotalHits = e.Body.SearchResults.TotalHits
	sr.Page = e.Body.SearchResults.Page
	sr.Language = e.Body.SearchResults.Language
	sr.Result = e.Body.SearchResults.Result

	return sr, nil
}
