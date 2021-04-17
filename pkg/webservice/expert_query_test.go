package webservice

import "testing"

func TestCreateExpertQuery(t *testing.T) {
	search := "samplesearch"

	q := NewExpertQuery(search)

	if q.String() != search {
		t.Errorf("Query was not set properly, got: %s, want: %s", q.String(), search)
	}
}
