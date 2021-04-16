package webservice

type SearchRequest struct {
	Username       string
	Password       string
	Query          *ExpertQuery
	Page           int64
	PageSize       int64
	SearchLanguage string
}

func NewSearchRequest(query string) *SearchRequest {
	q := NewExpertQuery(query)
	return &SearchRequest{
		Query: q,
	}
}

func NewSearchRequestWithConfig(cfg *Config, query string) *SearchRequest {
	r := NewSearchRequest(query)

	r.Username = cfg.Username
	r.Password = cfg.Password

	return r
}

func (s *SearchRequest) ToXML() string {
	return ""
}
