# eurlex-ws-go

# Overview

[eurlex-ws-go](https://github.com/slipke/eurlex-ws-go) provides an implementation of the [EUR-Lex](https://eur-lex.europa.eu/homepage.html) Webservice (Search) in Go. The API is described [here](https://eur-lex.europa.eu/content/help/webservice.html) and [here](https://eur-lex.europa.eu/content/tools/webservices/SearchWebServiceUserManual_v2.00.pdf).

## Install

```bash
go get -u github.com/slipke/eurlex-ws-go
```

## Usage

Create a simple query and call the API (replace `<username>` and `<password>` with your credentials):

```go
import (
    log "github.com/sirupsen/logrus"
	"github.com/slipke/eurlex-ws-go"
)

func main() {
	cfg := eurlex.NewConfig("<usename>", "<password>")
	ws := eurlex.NewWebservice(cfg)

	sr, err := ws.Search(
		eurlex.NewSearchRequestFromString("DN~32014R1338 OR DN~32014R1337"),
	)
	if err != nil {
		log.Fatalf("Failed to issue request: %s", err)
		return
	}

	log.Printf(
		"Returned %d elements (Page: %d, NumHits: %d, TotalHits: %d, Language: %s)",
		len(sr.Result),
		sr.Page,
		sr.NumHits,
		sr.TotalHits,
		sr.Language,
	)

	for _, r := range sr.Result {
		log.Printf("Result title: %s", r.Content.Notice.Expression.Title)
	}
}
```

## Config

eurlex.Config is used to configure your webservice, it has the following defaults:

```go
Config{
    ...
    Logger:   logrus.StandardLogger(),
    Client:   http.DefaultClient,
}
```

## SearchRequest

eurlex.SearchRequest is used to to query the API.

It has the following defaults:

```go
SearchRequest{
    Page:           1,
    PageSize:       10,
    SearchLanguage: "en",
}
```

`Page` defines the current page, whereas `PageSize` defines the maximum numbers of results in the response. `SearchLanguage` defines the language of the results.

## SearchResult

The structure of the returned `SearchResult` is derived from the returning XML. An example can be found in `fixtures/result.xml`. More information can be found under [Resources](#Resources).

## Todos

- Write test to verfiy XSD (see XSD 3)
- Improve code coverage


## Resources

- [EUR-Lex Help](https://eur-lex.europa.eu/content/help/webservice.html)
- [Web Service User Manual](https://eur-lex.europa.eu/content/tools/webservices/SearchWebServiceUserManual_v2.00.pdf)
- [Web Service Query Metadata](https://eur-lex.europa.eu/content/tools/webservices/WebServicesqueryMetadata.pdf)
- [Data Extraction Using Web Services](https://eur-lex.europa.eu/content/tools/webservices/DataExtractionUsingWebServices-v1.00.pdf)
- [WSDL](https://eur-lex.europa.eu/eurlex-ws?wsdl)
- [XSD 1](https://eur-lex.europa.eu/eurlex-ws?xsd=1)
- [XSD 2](https://eur-lex.europa.eu/eurlex-ws?xsd=2)
- [XSD 3](https://eur-lex.europa.eu/eurlex-ws?xsd=3)