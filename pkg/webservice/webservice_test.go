package webservice

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
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
	inputFile := "../../fixtures/result.xml"
	xmlBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("Failed to load fixtures %s: %s", inputFile, err)
	}

	cfg := NewConfig("testuser", "testpass")

	// Mock our HTTP Request
	cfg.Client = NewMockClient(func(req *http.Request) *http.Response {
		if req.URL.String() != wsURL {
			t.Errorf("Called wrong URL, got %s, want: %s", req.URL.String(), wsURL)
			t.FailNow()
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(string(xmlBytes))),
		}
	})

	ws := NewWebservice(cfg)

	sr, err := ws.Search(NewSearchRequestWithConfig(cfg, "testsearch"))
	if err != nil {
		t.Errorf("Failed to search: %s", err)
		t.FailNow()
	}

	if len(sr.Result) <= 0 {
		t.Errorf("Missing Results, got: %d, want: %s", len(sr.Result), "> 0")
	}
}
