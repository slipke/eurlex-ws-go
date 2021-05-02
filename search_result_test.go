package eurlex

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestInvalidSearchResult(t *testing.T) {
	er, err := NewSearchResultFromXML("")
	if er != nil {
		t.Errorf("Returned envelope where it should be nil")
	}

	if err == nil {
		t.Errorf("Should have returned an error message")
	}

	if !strings.Contains(err.Error(), errorFailedToCreateResult) {
		t.Errorf("Returned wrong error message")
	}
}

func TestCreateSearchResult(t *testing.T) {
	inputFile := "./fixtures/result.xml"
	xmlBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("Failed to load fixtures %s: %s", inputFile, err)
	}

	sr, err := NewSearchResultFromXML(string(xmlBytes))
	if err != nil {
		t.Errorf("Failed to create SearchResult: %s", err)
	}

	if sr == nil {
		t.Errorf("Faild to create SearchRequest, got nil")
	}

	if len(sr.Result) <= 0 {
		t.Errorf("Empty result set")
		t.FailNow()
	}

	checks := []struct {
		Label    string
		Value    interface{}
		Expected interface{}
	}{
		{"NumHits", sr.NumHits, int64(2)},
		{"TotalHits", sr.TotalHits, int64(2)},
		{"Page", sr.Page, int64(1)},
		{"Language", sr.Language, "en"},
		{"len(result)", len(sr.Result), 2},
	}

	for _, c := range checks {
		if c.Expected != c.Value {
			t.Errorf("Failed run Check '%s', want: %+v, have: %+v", c.Label, c.Expected, c.Value)
		}
	}

	// Check first result
	res := sr.Result[0]

	if len(res.Content.Notice.Expression) <= 0 {
		t.Errorf("Failed to check expressions, 0 found")
		t.FailNow()
	}
	e := res.Content.Notice.Expression[0]
	w := res.Content.Notice.Work

	if len(w.ResourceLegalPublishedInOfficialJournal) <= 0 {
		t.Errorf("Failed to check ResourceLegalPublishedInOfficialJournal, 0 found")
		t.FailNow()
	}
	rlpioj := w.ResourceLegalPublishedInOfficialJournal[0]

	if len(w.WorkCreatedByAgent) <= 0 {
		t.Errorf("Failed to check WorkCreatedByAgent, 0 found")
		t.FailNow()
	}
	wcba := w.WorkCreatedByAgent[0]

	checks = []struct {
		Label    string
		Value    interface{}
		Expected interface{}
	}{
		{"Result[0] Reference", res.Reference, "eng_cellar:2a2a817d-85b9-11e4-b8a5-01aa75ed71a1_en"},
		{"Result[0] Rank", res.Rank, "1"},
		{"Result[0] Notice Expression Title", e.ExpressionTitle, "Commission Implementing Regulation (EU) No 1337/2014 of 16 December 2014 amending Implementing Regulations (EU) No 947/2014 and (EU) No 948/2014 as regards the last day for submission of applications for private storage aid for butter and skimmed milk powder"},
		{"Result[0] Notice Expression Language", e.ExpressionUsesLanguage.URI.Identifier, "ENG"},
		{"Result[0] Notice Work IDCelex", w.IDCelex, "<em>32014R1337</em>"},
		{"Result[0] Notice Work ResourceLegalInForce", w.ResourceLegalInForce, false},
		{"Result[0] Notice Work ResourceLegalPublishedInOfficialJournal DatePublication", rlpioj.EmbeddedNotice.Work.DatePublication, "2014-12-17"},
		{"Result[0] Notice Work ResourceLegalPublishedInOfficialJournal len(SAMEAS)", len(rlpioj.SameAs), 3},
		{"Result[0] Notice Work WorkCreatedByAgent len(AltLAbel)", len(wcba.AltLabel), 4},
		{"Result[0] Notice Work WorkCreatedByAgent Identifier", wcba.Identifier, "COM"},
		{"Result[0] Notice Work WorkDateDocument Day", w.WorkDateDocument.Day, "16"},
		{"Result[0] Notice Work len(HasExpression)", len(w.WorkHasExpression), 23},
	}

	for _, c := range checks {
		if c.Expected != c.Value {
			t.Errorf("Failed run Check '%s', want: %+v, have: %+v", c.Label, c.Expected, c.Value)
		}
	}

	if len(w.WorkHasExpression) <= 0 {
		t.Errorf("Failed to check WorkHasExpression, 0 found")
		t.FailNow()
	}
	// Check the first WorkHasExpression
	whe := w.WorkHasExpression[0]

	if len(whe.EmbeddedNotice.Expression) <= 0 {
		t.Errorf("Failed to check whe.EmbeddedNotice.Expression, 0 found")
		t.FailNow()
	}
	// Get inner expression
	ie := whe.EmbeddedNotice.Expression[0]

	if len(whe.EmbeddedNotice.Manifestation) <= 0 {
		t.Errorf("Failed to check whe.EmbeddedNotice.Manifestation, 0 found")
		t.FailNow()
	}
	// Get inner manifestation
	im := whe.EmbeddedNotice.Manifestation[0]

	checks = []struct {
		Label    string
		Value    interface{}
		Expected interface{}
	}{
		{"Expression[0] EmbeddedNotice Expression Language", ie.ExpressionUsesLanguage.OpCode, "ENG"},
		{"Expression[0] EmbeddedNotice Manifestation PageFirst", im.ManifestationOfficialJournalPartPageFirst, "0015"},
	}

	for _, c := range checks {
		if c.Expected != c.Value {
			t.Errorf("Failed run Check '%s', want: %+v, have: %+v", c.Label, c.Expected, c.Value)
		}
	}
}

func TestMissingBody(t *testing.T) {
	inputFile := "./fixtures/result_empty_body.xml"
	xmlBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("Failed to load fixtures %s: %s", inputFile, err)
	}

	sr, err := NewSearchResultFromXML(string(xmlBytes))
	if err == nil {
		t.Errorf("Expected error message due to missing body")
	}

	if sr != nil {
		t.Errorf("Expected ErrorResponse to be nil, due to missing body element")
	}
}

func TestMissingSearchResults(t *testing.T) {
	inputFile := "./fixtures/result_empty_searchresult.xml"
	xmlBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("Failed to load fixtures %s: %s", inputFile, err)
	}

	sr, err := NewSearchResultFromXML(string(xmlBytes))
	if err == nil {
		t.Errorf("Expected error message due to missing SearchResult")
	}

	if sr != nil {
		t.Errorf("Expected ErrorResponse to be nil, due to missing SearchResult element")
	}
}
