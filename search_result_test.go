package eurlex

import (
	"io/ioutil"
	"testing"
)

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

	checks = []struct {
		Label    string
		Value    interface{}
		Expected interface{}
	}{
		{"Result[0] Reference", res.Reference, "eng_cellar:2a2a817d-85b9-11e4-b8a5-01aa75ed71a1_en"},
		{"Result[0] Rank", res.Rank, "1"},
		{"Result[0] Notice Expression Title", res.Content.Notice.Expression.Title, "Commission Implementing Regulation (EU) No 1337/2014 of 16 December 2014 amending Implementing Regulations (EU) No 947/2014 and (EU) No 948/2014 as regards the last day for submission of applications for private storage aid for butter and skimmed milk powder"},
		{"Result[0] Notice Expression Language", res.Content.Notice.Expression.Language.URI.Identifier, "ENG"},
		{"Result[0] Notice Work IDCelex", res.Content.Notice.Work.IDCelex, "<em>32014R1337</em>"},
		{"Result[0] Notice Work ResourceLegalInForce", res.Content.Notice.Work.ResourceLegalInForce, false},
		{"Result[0] Notice Work ResourceLegalPublishedInOfficialJournal DatePUblication", res.Content.Notice.Work.ResourceLegalPublishedInOfficialJournal.EmbeddedNotice.Work.DatePublication, "2014-12-17"},
		{"Result[0] Notice Work ResourceLegalPublishedInOfficialJournal len(SAMEAS)", len(res.Content.Notice.Work.ResourceLegalPublishedInOfficialJournal.Samas), 3},
		{"Result[0] Notice Work WorkCreatedByAgent len(AltLAbel)", len(res.Content.Notice.Work.WorkCreatedByAgent.AltLabel), 4},
		{"Result[0] Notice Work WorkCreatedByAgent Identifier", res.Content.Notice.Work.WorkCreatedByAgent.Identifier, "COM"},
		{"Result[0] Notice Work WorkDateDocument Day", res.Content.Notice.Work.WorkDateDocument.Day, "16"},
		{"Result[0] Notice Work len(HasExpression)", len(res.Content.Notice.Work.WorkHasExpression), 23},
	}

	for _, c := range checks {
		if c.Expected != c.Value {
			t.Errorf("Failed run Check '%s', want: %+v, have: %+v", c.Label, c.Expected, c.Value)
		}
	}

	// Check the first expression
	exp := res.Content.Notice.Work.WorkHasExpression[0]

	checks = []struct {
		Label    string
		Value    interface{}
		Expected interface{}
	}{
		{"Expression[0] EmbeddedNotice Expression Language", exp.EmbeddedNotice.Expression.ExpressionUsesLanguage, "ENG"},
		{"Expression[0] EmbeddedNotice Manifestation PageFirst", exp.EmbeddedNotice.Manifestation.ManifestationOfficialJournalPartPageFirst, "0015"},
	}

	for _, c := range checks {
		if c.Expected != c.Value {
			t.Errorf("Failed run Check '%s', want: %+v, have: %+v", c.Label, c.Expected, c.Value)
		}
	}
}
