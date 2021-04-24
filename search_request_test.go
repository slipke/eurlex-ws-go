package eurlex

import (
	"strings"
	"testing"
)

func TestCreateSearchRequest(t *testing.T) {
	search := "testsearch"

	sr := NewSearchRequestFromString(search)

	if sr.Query.String() != search {
		t.Errorf("Query was not set properly, got: %s, want: %s", sr.Query.String(), search)
	}

	sr = NewSearchRequestFromString(search)

	if sr.Query.String() != search {
		t.Errorf("Query was not set properly, got: %s, want: %s", sr.Query.String(), search)
	}
}

func TestToXML(t *testing.T) {
	search := "testsearch"

	sr := NewSearchRequestFromString(search)
	sr.Page = 1
	sr.PageSize = 20
	sr.SearchLanguage = "de"
	sr.username = "testuser"
	sr.password = "testpass"

	requiredElements := []string{
		// Header
		"wsse:Security",
		"xmlns:wsse=\"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd\"",
		"soap:mustUnderstand=\"true\"",
		"wsse:UsernameToken",
		"xmlns:wsu=\"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd\"",
		"wsu:Id=\"UsernameToken-1\"",
		"wsse:Username",
		"wsse:Password",
		"Type=\"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText\"",
		"/wsse:Username",
		"/wsse:Password",
		"testuser",
		"testpass",
		"/wsse:UsernameToken",
		"/wsse:Security",
		// Body
		"sear:searchRequest",
		"sear:expertQuery",
		"<sear:expertQuery><![CDATA[testsearch]]></sear:expertQuery>",
		"sear:page",
		"/sear:page",
		"sear:pageSize",
		"/sear:pageSize",
		"sear:searchLanguage",
		"/sear:searchLanguage",
		"<sear:page>1</sear:page>",
		"<sear:pageSize>20</sear:pageSize>",
		"<sear:searchLanguage>de</sear:searchLanguage>",
	}

	isXML, err := sr.ToXML()
	if err != nil {
		t.Errorf("ToXML failed: %s", err)
	}

	for _, e := range requiredElements {
		if !strings.Contains(string(isXML), e) {
			t.Errorf("Element %s was not found in resulting XML", e)
		}
	}
}
