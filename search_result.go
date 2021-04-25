package eurlex

import (
	"fmt"

	"github.com/slipke/eurlex-ws-go/internal/response"
)

const (
	errorFailedToCreateResult = "failed to create envelope from XML"
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
		return nil, fmt.Errorf("%s: %s", errorFailedToCreateResult, err)
	}

	if e.Body == nil {
		return nil, fmt.Errorf("failed to parse search result, missing body")
	}

	if e.Body.SearchResults == nil {
		return nil, fmt.Errorf("failed to parse search result, missing SearchResults element")
	}

	sr.NumHits = e.Body.SearchResults.NumHits
	sr.TotalHits = e.Body.SearchResults.TotalHits
	sr.Page = e.Body.SearchResults.Page
	sr.Language = e.Body.SearchResults.Language
	sr.Result = e.Body.SearchResults.Result

	return sr, nil
}
