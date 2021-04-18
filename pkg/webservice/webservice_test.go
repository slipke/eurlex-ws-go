package webservice

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	// @TODO Replace and adjust tests
	xmlWebservice = `<S:Envelope xmlns:S="http://www.w3.org/2003/05/soap-envelope">
    <S:Body>
        <searchResults xsi:schemaLocation="http://eur-lex.europa.eu/search http://localhost:7001/eurlex-frontoffice/eurlex-ws?xsd=3"
            xmlns="http://eurlex.europa.eu/search"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
            <numhits>10</numhits>
            <totalhits>1946</totalhits>
            <page>1</page>
            <language>en</language>
            <result>
                <reference>eng_cellar:93836665-712f-4444-a1e6-dadad5607e80_en</reference>
                <rank>1</rank>
                <content>
                    <NOTICE>
                        <EXPRESSION>
                            <EXPRESSION_TITLE>
                                <VALUE>Decision on the …</VALUE>
                            </EXPRESSION_TITLE>
                            <EXPRESSION_USES_LANGUAGE>
                                <URI>
                                    <IDENTIFIER>ENG</IDENTIFIER>
                                </URI>
                            </EXPRESSION_USES_LANGUAGE>
                        </EXPRESSION>
                    </NOTICE>
                </content>
            </result>
        </searchResults>
    </S:Body>
</S:Envelope>`
)

func TestCreateWebservice(t *testing.T) {
	u := "testuser"
	p := "testpass"

	cfg := NewConfig(u, p)

	ws := NewWebservice(cfg)

	if ws.cfg != cfg {
		t.Errorf("Config not set properly, got: %+v, want: %+v", ws.cfg, cfg)
	}
}

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewMockClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestSearch(t *testing.T) {
	cfg := NewConfig("testuser", "testpass")

	// Mock our HTTP Request
	cfg.Client = NewMockClient(func(req *http.Request) *http.Response {
		if req.URL.String() != wsURL {
			t.Errorf("Called wrong URL, got %s, want: %s", req.URL.String(), wsURL)
			t.FailNow()
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(xmlWebservice)),
		}
	})

	ws := NewWebservice(cfg)

	sr, err := ws.Search(NewSearchRequestWithConfig(cfg, "testsearch"))
	if err != nil {
		t.Errorf("Failed to search: %s", err)
		t.FailNow()
	}

	if sr.NumHits != 10 {
		t.Errorf("Failed to assign NumHits, got: %d, want: %d", sr.NumHits, 10)
	}

	if sr.TotalHits != 1946 {
		t.Errorf("Failed to assign TotalHits, got: %d, want: %d", sr.TotalHits, 1946)
	}

	if sr.Page != 1 {
		t.Errorf("Failed to assign Page, got: %d, want: %d", sr.Page, 1)
	}

	if sr.Language != "en" {
		t.Errorf("Failed to assign Language, got: %s, want: %s", sr.Language, "en")
	}

	// @TODO More tests for result
	if sr.Result == nil {
		t.Errorf("Failed to assign Result, got: %+v, want: %+v", sr.Result, "not nil")
	}
}
