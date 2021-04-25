package response

import (
	"io/ioutil"
	"testing"
)

func TestInvalidErrorResponse(t *testing.T) {
	er, err := NewEnvelopeFromXML("")
	if er != nil {
		t.Errorf("Returned envelope where it should be nil")
	}

	if err == nil {
		t.Errorf("Should have returned an error message")
	}
}

func TestSuccessResponse(t *testing.T) {
	inputFile := "../../fixtures/result.xml"
	xmlBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("Failed to load fixtures %s: %s", inputFile, err)
	}

	e, err := NewEnvelopeFromXML(string(xmlBytes))
	if err != nil {
		t.Errorf("Failed to create SearchResult: %s", err)
	}

	if e.Body == nil {
		t.Errorf("Body not set, got: %+v", e.Body)
	}
}

func TestErrorResponseAuthentication(t *testing.T) {
	inputFile := "../../fixtures/error_response_authentication.xml"
	xmlBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("Failed to load fixtures %s: %s", inputFile, err)
	}

	e, err := NewEnvelopeFromXML(string(xmlBytes))
	if err != nil {
		t.Errorf("NewEnvelopeFromXML failed: %s", err)
	}

	if e.Header == nil {
		t.Errorf("Header not set, got: %+v", e.Header)
	}

	if e.Body == nil {
		t.Errorf("Body not set, got: %+v", e.Body)
	}

	if e.Body.SearchResults != nil {
		t.Errorf("Found SearchResults which should not be there, got: %+v, want: %+v", e.Body.SearchResults, nil)
	}

	er := e.Body.Fault

	if er == nil {
		t.Errorf("ErrorResponse not set, got %+v", er)
	}

	if er.Code != "env:Sender" {
		t.Errorf("Code not set, got %s, want: %s", er.Code, "env:Sender")
	}

	if er.Reason != "Failed to assert identity with UsernameToken." {
		t.Errorf("Reason not set, got %s, want: %s", er.Reason, "Failed to assert identity with UsernameToken.")
	}
}
