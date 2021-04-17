package webservice

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// wsdlURL = "https://eur-lex.europa.eu/eurlex-ws?wsdl"
	// @TODO Copied from current WSDL above, parse automatically for each request?
	wsURL = "https://eur-lex.europa.eu/EURLexWebService"
)

type WebserviceInterface interface {
	Search(sr *SearchRequest) (*SearchResult, error)
}

type Webservice struct {
	cfg *Config
}

func (ws *Webservice) Search(sr *SearchRequest) (*SearchResult, error) {
	xml, err := sr.ToXML()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SearchRequest: %s", err)
	}

	body := bytes.NewReader(xml)
	r, err := http.NewRequest(http.MethodPost, wsURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}

	res, err := ws.cfg.Client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %s", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	// @TODO how to differentiate error responses from valid ones? (HTTP header?)
	// => New different parser / object?
	sRes, err := NewSearchResultFromXML(string(resBody))
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %s", err)
	}

	return sRes, nil
}

func NewWebservice(cfg *Config) *Webservice {
	return &Webservice{
		cfg: cfg,
	}
}
