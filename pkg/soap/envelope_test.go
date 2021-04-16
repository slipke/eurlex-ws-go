package soap

import (
	"testing"
)

func TestCreateEnvelope(t *testing.T) {
	e := NewEnvelope()

	// This is how our xml should look like
	shouldXML := `<soap:Envelope xmlns:sear="http://eur-lex.europa.eu/search" xmlns:soap="http://www.w3.org/2003/05/soap-envelope"><soap:Header><wsse:Security xmlns:wsse="http://docs.oasisopen.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" soap:mustUnderstand="true"><wsse:UsernameToken wsu:Id="UsernameToken-3" xmlns:wsu="http://docs.oasisopen.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"><wsse:Username></wsse:Username><wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"></wsse:Password></wsse:UsernameToken></wsse:Security></soap:Header><soap:Body></soap:Body></soap:Envelope>`

	isXML, err := e.ToXML()
	if err != nil {
		t.Errorf("ToXML failed: %s", err)
	}

	if isXML != shouldXML {
		t.Errorf("XML output wrong, got: %s, want: %s", isXML, shouldXML)
	}
}

func TestHeader(t *testing.T) {
	e := NewEnvelope()
	e.Header.Security.UsernameToken.Username = "testuser"
	e.Header.Security.UsernameToken.Password.Password = "testpass"

	// This is how our xml should look like
	shouldXML := `<soap:Envelope xmlns:sear="http://eur-lex.europa.eu/search" xmlns:soap="http://www.w3.org/2003/05/soap-envelope"><soap:Header><wsse:Security xmlns:wsse="http://docs.oasisopen.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" soap:mustUnderstand="true"><wsse:UsernameToken wsu:Id="UsernameToken-3" xmlns:wsu="http://docs.oasisopen.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"><wsse:Username>testuser</wsse:Username><wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">testpass</wsse:Password></wsse:UsernameToken></wsse:Security></soap:Header><soap:Body></soap:Body></soap:Envelope>`

	isXML, err := e.ToXML()
	if err != nil {
		t.Errorf("ToXML failed: %s", err)
	}

	if isXML != shouldXML {
		t.Errorf("XML output wrong, got: %s, want: %s", isXML, shouldXML)
	}
}

func TestBody(t *testing.T) {
	e := NewEnvelope()
	e.Header.Security.UsernameToken.Username = "testuser"
	e.Header.Security.UsernameToken.Password.Password = "testpass"

	e.Body.RootElement = NewSearchRequest("testsearch", 1, 10, "de")

	// This is how our xml should look like
	shouldXML := `<soap:Envelope xmlns:sear="http://eur-lex.europa.eu/search" xmlns:soap="http://www.w3.org/2003/05/soap-envelope"><soap:Header><wsse:Security xmlns:wsse="http://docs.oasisopen.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" soap:mustUnderstand="true"><wsse:UsernameToken wsu:Id="UsernameToken-3" xmlns:wsu="http://docs.oasisopen.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"><wsse:Username>testuser</wsse:Username><wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">testpass</wsse:Password></wsse:UsernameToken></wsse:Security></soap:Header><soap:Body><sear:searchRequest><sear:expertQuery>testsearch</sear:expertQuery><sear:page>1</sear:page><sear:pageSize>10</sear:pageSize><sear:searchLanguage>de</sear:searchLanguage></sear:searchRequest></soap:Body></soap:Envelope>`

	isXML, err := e.ToXML()
	if err != nil {
		t.Errorf("ToXML failed: %s", err)
	}

	if isXML != shouldXML {
		t.Errorf("XML output wrong, got: %s, want: %s", isXML, shouldXML)
	}
}
