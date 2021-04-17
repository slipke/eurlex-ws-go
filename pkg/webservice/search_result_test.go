package webservice

import "testing"

const (
	xmlResult = `<S:Envelope xmlns:S="http://www.w3.org/2003/05/soap-envelope">
    <S:Body>
        <searchResults xsi:schemaLocation="http://eur-lex.europa.eu/search http://localhost:7001/eurlex-frontoffice/eurlex-ws?xsd=3"
            xmlns="http://eurlex.europa.eu/search"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
            <numhits>10</numhits>
            <totalhits>1946</totalhits>
            <page>1</page>
            <language>en</language>
            <result>
                <reference>eng_cellar:93836665-712f-4444-a1e6-dadad5607e80_en</reference>
                <rank>1</rank>
                <content>
                    <NOTICE>
                        <EXPRESSION>
                            <EXPRESSION_TITLE>
                                <VALUE>Decision on the …</VALUE>
                            </EXPRESSION_TITLE>
                            <EXPRESSION_USES_LANGUAGE>
                                <URI>
                                    <IDENTIFIER>ENG</IDENTIFIER>
                                </URI>
                            </EXPRESSION_USES_LANGUAGE>
                        </EXPRESSION>
                    </NOTICE>
                </content>
            </result>
        </searchResults>
    </S:Body>
</S:Envelope>`
)

func TestCreateSearchResult(t *testing.T) {
	sr, err := NewSearchResultFromXML(xmlResult)
	if err != nil {
		t.Errorf("Failed to create SearchResult: %s", err)
	}

	if sr == nil {
		t.Errorf("Faild to create SearchRequest, got nil")
	}

	if sr.NumHits != 10 {
		t.Errorf("Failed to assign NumHits, got: %d, want: %d", sr.NumHits, 10)
	}

	if sr.TotalHits != 1946 {
		t.Errorf("Failed to assign TotalHits, got: %d, want: %d", sr.TotalHits, 1946)
	}

	if sr.Page != 1 {
		t.Errorf("Failed to assign Page, got: %d, want: %d", sr.Page, 1)
	}

	if sr.Language != "en" {
		t.Errorf("Failed to assign Language, got: %s, want: %s", sr.Language, "en")
	}

	if sr.Result == nil {
		t.Errorf("Failed to assign Result, got: %+v, want: %+v", sr.Result, "not nil")
	}

	// @TODO Add more tests on result
}
