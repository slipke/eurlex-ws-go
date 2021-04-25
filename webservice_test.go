package eurlex

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
	inputFile := "./fixtures/result.xml"
	xmlBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("Failed to load fixtures %s: %s", inputFile, err)
		t.FailNow()
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

	sr, err := ws.Search(NewSearchRequestFromString("testsearch"))
	if err != nil {
		t.Errorf("Failed to search: %s", err)
		t.FailNow()
	}

	if len(sr.Result) <= 0 {
		t.Errorf("Missing Results, got: %d, want: %s", len(sr.Result), "> 0")
	}
}

func TestInvalidConfig(t *testing.T) {
	cfg := NewConfig("", "")

	// Mock our HTTP Request
	cfg.Client = NewMockClient(func(req *http.Request) *http.Response {
		if req.URL.String() != wsURL {
			t.Errorf("Called wrong URL, got %s, want: %s", req.URL.String(), wsURL)
			t.FailNow()
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader("")),
		}
	})

	ws := NewWebservice(cfg)

	sr, err := ws.Search(NewSearchRequestFromString("testsearch"))
	if err == nil {
		t.Errorf("Should return an error messsage due to empty username or password")
	}

	if sr != nil {
		t.Errorf("SearchResult expected to be nil, due to invalid config")
	}
}

func TestEmptySearchRequest(t *testing.T) {
	ws := NewWebservice(NewConfig("testuser", "testpass"))
	_, err := ws.Search(nil)
	if err == nil {
		t.Errorf("Expected error message due to empty SearchRequest")
	}
}

func TestInvalidResponse(t *testing.T) {
	cfg := NewConfig("testuser", "testpass")

	// Mock our HTTP Request
	cfg.Client = NewMockClient(func(req *http.Request) *http.Response {
		if req.URL.String() != wsURL {
			t.Errorf("Called wrong URL, got %s, want: %s", req.URL.String(), wsURL)
			t.FailNow()
		}
		return nil
	})

	ws := NewWebservice(cfg)

	sr, err := ws.Search(NewSearchRequestFromString("testsearch"))
	if sr != nil {
		t.Errorf("Expected result to be nil, due to invalid response from HTTP request")
	}

	if err == nil {
		t.Errorf("Expected error due to invalid HTTP response")
	}
}

func TestEmptyResponseBody(t *testing.T) {
	cfg := NewConfig("testuser", "testpass")

	// Mock our HTTP Request
	cfg.Client = NewMockClient(func(req *http.Request) *http.Response {
		if req.URL.String() != wsURL {
			t.Errorf("Called wrong URL, got %s, want: %s", req.URL.String(), wsURL)
			t.FailNow()
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader("")),
		}
	})

	ws := NewWebservice(cfg)

	sr, err := ws.Search(NewSearchRequestFromString("testsearch"))
	if sr != nil {
		t.Errorf("Expected result to be nil, due to invalid response Body")
	}

	if err == nil {
		t.Errorf("Expected error due to invalid HTTP response Body")
	}
}

func TestInvalidResponseBody(t *testing.T) {
	cfg := NewConfig("testuser", "testpass")

	// Mock our HTTP Request
	cfg.Client = NewMockClient(func(req *http.Request) *http.Response {
		if req.URL.String() != wsURL {
			t.Errorf("Called wrong URL, got %s, want: %s", req.URL.String(), wsURL)
			t.FailNow()
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader("{test:'asdf'}")),
		}
	})

	ws := NewWebservice(cfg)

	sr, err := ws.Search(NewSearchRequestFromString("testsearch"))
	if sr != nil {
		t.Errorf("Expected result to be nil, due to invalid response Body")
	}

	if err == nil {
		t.Errorf("Expected error due to invalid HTTP response Body")
	}
}

func TestInvalidSearchQuery(t *testing.T) {
	inputFile := "./fixtures/error_response_invalid_query.xml"
	xmlBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("Failed to load fixtures %s: %s", inputFile, err)
		t.FailNow()
	}

	cfg := NewConfig("testuser", "testpass")

	// Mock our HTTP Request
	cfg.Client = NewMockClient(func(req *http.Request) *http.Response {
		if req.URL.String() != wsURL {
			t.Errorf("Called wrong URL, got %s, want: %s", req.URL.String(), wsURL)
			t.FailNow()
		}
		return &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       ioutil.NopCloser(strings.NewReader(string(xmlBytes))),
		}
	})

	ws := NewWebservice(cfg)

	sr, err := ws.Search(NewSearchRequestFromString("asdf asdf"))
	if sr != nil {
		t.Errorf("Expected result to be nil, due to invalid query")
	}

	if err == nil {
		t.Errorf("Expected error due to invalid query")
	}
}
