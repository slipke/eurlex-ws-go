package webservice

import "testing"

func TestCreateSearchRequest(t *testing.T) {
	search := "testsearch"

	sr := NewSearchRequest(search)

	if sr.Query.String() != search {
		t.Errorf("Query was not set properly, got: %s, want: %s", sr.Query.String(), search)
	}

	username := "testuser"
	password := "testpass"

	cfg := NewConfig(username, password)
	sr = NewSearchRequestWithConfig(cfg, search)

	if sr.Username != username {
		t.Errorf("Username was not set properly, got: %s, want: %s", sr.Username, username)
	}

	if sr.Password != password {
		t.Errorf("Password was not set properly, got: %s, want: %s", sr.Password, password)
	}

	if sr.Query.String() != search {
		t.Errorf("Query was not set properly, got: %s, want: %s", sr.Query.String(), search)
	}

	// @TODO Add more fields when decided
}

func TestToXML(t *testing.T) {
	search := "testsearch"
	u := "testuser"
	p := "testpass"

	sr := NewSearchRequestWithConfig(NewConfig(u, p), search)
	sr.Page = 1
	sr.PageSize = 10
	sr.SearchLanguage = "de"

	// This is how our xml should look like
	shouldXML := `<soap:Envelope xmlns:sear="http://eur-lex.europa.eu/search" xmlns:soap="http://www.w3.org/2003/05/soap-envelope"><soap:Header><wsse:Security xmlns:wsse="http://docs.oasisopen.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" soap:mustUnderstand="true"><wsse:UsernameToken wsu:Id="UsernameToken-3" xmlns:wsu="http://docs.oasisopen.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"><wsse:Username>testuser</wsse:Username><wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">testpass</wsse:Password></wsse:UsernameToken></wsse:Security></soap:Header><soap:Body><sear:searchRequest><sear:expertQuery>testsearch</sear:expertQuery><sear:page>1</sear:page><sear:pageSize>10</sear:pageSize><sear:searchLanguage>de</sear:searchLanguage></sear:searchRequest></soap:Body></soap:Envelope>`

	isXML, err := sr.ToXML()
	if err != nil {
		t.Errorf("ToXML failed: %s", err)
	}

	if string(isXML) != shouldXML {
		t.Errorf("XML output wrong, got: %s, want: %s", isXML, shouldXML)
	}

}
