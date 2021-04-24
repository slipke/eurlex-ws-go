package eurlex

type ExpertQuery struct {
	// @TODO Currently, only s
	// Later this contains multiple fields to construct a complex query
	s string
}

func NewExpertQueryFromString(query string) *ExpertQuery {
	return &ExpertQuery{
		s: query,
	}
}

func (e *ExpertQuery) String() string {
	return e.s
}
