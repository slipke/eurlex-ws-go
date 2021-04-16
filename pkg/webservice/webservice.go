package webservice

const (
	wsdlURL = "https://eur-lex.europa.eu/eurlex-ws?wsdl"
)

type WebserviceInterface interface {
	Search(sr *SearchRequest) (*SearchResult, error)
}

type Webservice struct {
	cfg *Config
}

func NewWebservice(cfg *Config) *Webservice {
	return &Webservice{
		cfg: cfg,
	}
}
