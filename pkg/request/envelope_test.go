package request

import (
	"strings"
	"testing"
)

func TestCreateEnvelope(t *testing.T) {
	e := NewEnvelope()

	requiredElements := []string{
		"soap:Envelope",
		"xmlns:soap=\"http://www.w3.org/2003/05/soap-envelope\"",
		"xmlns:sear=\"http://eur-lex.europa.eu/search\"",
		"soap:Header",
		"/soap:Header",
		"soap:Body",
		"/soap:Body",
		"/soap:Envelope",
	}

	isXML, err := e.ToXML()
	if err != nil {
		t.Errorf("ToXML failed: %s", err)
	}

	for _, e := range requiredElements {
		if !strings.Contains(string(isXML), e) {
			t.Errorf("Element %s was not found in resulting XML", e)
		}
	}
}

func TestHeader(t *testing.T) {
	e := NewEnvelope()
	e.Header.RootElement = NewSecurity("testuser", "testpass")

	requiredElements := []string{
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
	}

	isXML, err := e.ToXML()
	if err != nil {
		t.Errorf("ToXML failed: %s", err)
	}

	for _, e := range requiredElements {
		if !strings.Contains(string(isXML), e) {
			t.Errorf("Element %s was not found in resulting XML", e)
		}
	}
}

func TestBody(t *testing.T) {
	e := NewEnvelope()
	e.Body.RootElement = NewSearchRequest("testsearch", 1, 10, "de")

	requiredElements := []string{
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
		"<sear:pageSize>10</sear:pageSize>",
		"<sear:searchLanguage>de</sear:searchLanguage>",
	}

	isXML, err := e.ToXML()
	if err != nil {
		t.Errorf("ToXML failed: %s", err)
	}

	for _, e := range requiredElements {
		if !strings.Contains(string(isXML), e) {
			t.Errorf("Element %s was not found in resulting XML", e)
		}
	}
}
