package webservice

import (
	"io/ioutil"
	"testing"
)

func TestAuthenticationErrorResponseFromXML(t *testing.T) {
	inputFile := "../../fixtures/error_response_authentication.xml"
	xmlBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("Failed to load fixtures %s: %s", inputFile, err)
	}

	er, err := NewErrorResponseFromXML(string(xmlBytes))
	if err != nil {
		t.Errorf("Failed to unmarshal error response from XML: %s", err)
		t.FailNow()
	}

	if er.Code != "env:Sender" {
		t.Errorf("Invalid field Code, got %s, want %s", er.Code, "env:Sender")
	}

	if er.Reason != "Failed to assert identity with UsernameToken." {
		t.Errorf("Invalid field Reason, got %s, want: %s", er.Reason, "Failed to assert identity with UsernameToken.")
	}

	// @TODO Add more tests for different error responses (i.e. wrong query)
}
