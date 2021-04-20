package response

import (
	"encoding/xml"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  *Header  `xml:"Header"`
	Body    *Body    `xml:"Body"`
}

func NewEnvelopeFromXML(xmlStr string) (*Envelope, error) {
	e := &Envelope{}
	if err := xml.Unmarshal([]byte(xmlStr), e); err != nil {
		return nil, err
	}
	return e, nil
}

type Header struct {
	XMLName       xml.Name       `xml:"Header"`
	NotUnderstood *NotUnderstood `xml:"NotUnderstood"`
}

type NotUnderstood struct {
	XMLName xml.Name `xml:"NotUnderstood"`
	Qname   string   `xml:"qname,attr"`
}

type Body struct {
	XMLName       xml.Name       `xml:"Body"`
	Fault         *Fault         `xml:"Fault"`
	SearchResults *SearchResults `xml:"searchResults"`
}

type Fault struct {
	XMLName xml.Name `xml:"Fault"`
	Code    string   `xml:"Code>Value"`
	Reason  string   `xml:"Reason>Text"`
}

type SearchResults struct {
	XMLName   xml.Name  `xml:"searchResults"`
	NumHits   int64     `xml:"numhits"`
	TotalHits int64     `xml:"totalhits"`
	Page      int64     `xml:"page"`
	Language  string    `xml:"language"`
	Result    []*Result `xml:"result"`
}

type Result struct {
	XMLName   xml.Name `xml:"result"`
	Reference string   `xml:"reference"`
	Rank      string   `xml:"rank"`
	Content   *Content `xml:"content"`
}

type Content struct {
	XMLName xml.Name `xml:"content"`
	Notice  *Notice  `xml:"NOTICE"`
}

type Notice struct {
	XMLName    xml.Name    `xml:"NOTICE"`
	Expression *Expression `xml:"EXPRESSION"`
	Work       *NoticeWork `xml:"WORK"`
}

type Expression struct {
	XMLName  xml.Name                `xml:"EXPRESSION"`
	Title    string                  `xml:"EXPRESSION_TITLE>VALUE"`
	Language *ExpressionUsesLanguage `xml:"EXPRESSION_USES_LANGUAGE"`
}

type ExpressionUsesLanguage struct {
	XMLName xml.Name `xml:"EXPRESSION_USES_LANGUAGE"`
	URI     *URI     `xml:"URI"`
}

type NoticeWork struct {
	XMLName                                 xml.Name                                 `xml:"WORK"`
	IDCelex                                 string                                   `xml:"ID_CELEX>VALUE"`
	ResourceLegalInForce                    bool                                     `xml:"RESOURCE_LEGAL_IN-FORCE>VALUE"`
	ResourceLegalPublishedInOfficialJournal *ResourceLegalPublishedInOfficialJournal `xml:"RESOURCE_LEGAL_PUBLISHED_IN_OFFICIAL-JOURNAL"`
	WorkCreatedByAgent                      *WorkCreatedByAgent                      `xml:"WORK_CREATED_BY_AGENT"`
	WorkDateDocument                        *WorkDateDocument                        `xml:"WORK_DATE_DOCUMENT"`
	WorkHasExpression                       []*WorkHasExpression                     `xml:"WORK_HAS_EXPRESSION"`
}

type ResourceLegalPublishedInOfficialJournal struct {
	XMLName        xml.Name        `xml:"RESOURCE_LEGAL_PUBLISHED_IN_OFFICIAL-JOURNAL"`
	EmbeddedNotice *EmbeddedNotice `xml:"EMBEDDED_NOTICE"`
	Samas          []*Sameas       `xml:"SAMEAS"`
}

type EmbeddedNotice struct {
	XMLName xml.Name            `xml:"EMBEDDED_NOTICE"`
	Work    *EmbeddedNoticeWork `xml:"WORK"`
}

type EmbeddedNoticeWork struct {
	XMLName                                 xml.Name                                 `xml:"WORK"`
	DatePublication                         string                                   `xml:"DATE_PUBLICATION>VALUE"`
	OfficialJournalClass                    string                                   `xml:"OFFICIAL-JOURNAL_CLASS>VALUE"`
	OfficialJournalNumber                   string                                   `xml:"OFFICIAL-JOURNAL_NUMBER>VALUE"`
	OfficialJournalPartOfCollectionDocument *OfficialJournalPartOfCollectionDocument `xml:"OFFICIAL-JOURNAL_PART_OF_COLLECTION_DOCUMENT"`
	OfficialJournalYear                     string                                   `xml:"OFFICIAL-JOURNAL_YEAR>VALUE"`
}

type OfficialJournalPartOfCollectionDocument struct {
	XMLName xml.Name `xml:"OFFICIAL-JOURNAL_PART_OF_COLLECTION_DOCUMENT"`
	URI     *URI     `xml:"URI"`
}

type Sameas struct {
	XMLName xml.Name `xml:"SAMEAS"`
	URI     *URI     `xml:"URI"`
}

type URI struct {
	XMLName    xml.Name `xml:"URI"`
	Identifier string   `xml:"IDENTIFIER"`
	Type       string   `xml:"TYPE"`
	Value      string   `xml:"VALUE"`
}

type WorkCreatedByAgent struct {
	XMLName    xml.Name `xml:"WORK_CREATED_BY_AGENT"`
	AltLabel   []string `xml:"ALTLABEL"`
	CompactURI string   `xml:"COMPACT_URI"`
	Identifier string   `xml:"IDENTIFIER"`
	OpCode     string   `xml:"OP-CODE"`
	PrefLabel  string   `xml:"PREFLABEL"`
	URI        *URI     `xml:"URI"`
}

type WorkDateDocument struct {
	XMLName xml.Name `xml:"WORK_DATE_DOCUMENT"`
	Day     string   `xml:"DAY"`
	Month   string   `xml:"MONTH"`
	Year    string   `xml:"YEAR"`
}

type WorkHasExpression struct {
	XMLName        xml.Name                         `xml:"WORK_HAS_EXPRESSION"`
	EmbeddedNotice *WorkHasExpressionEmbeddedNotice `xml:"EMBEDDED_NOTICE"`
}

type WorkHasExpressionEmbeddedNotice struct {
	XMLName       xml.Name                     `xml:"EMBEDDED_NOTICE"`
	Expression    *WorkHasExpressionExpression `xml:"EXPRESSION"`
	Manifestation *Manifestation               `xml:"MANIFESTATION"`
}

type WorkHasExpressionExpression struct {
	XMLName                xml.Name `xml:"EXPRESSION"`
	ExpressionUsesLanguage string   `xml:"EXPRESSION_USES_LANGUAGE>OP-CODE"`
}

type Manifestation struct {
	XMLName                                   xml.Name                          `xml:"MANIFESTATION"`
	ManifestationOfficialJournalPartPageFirst string                            `xml:"MANIFESTATION_OFFICIAL-JOURNAL_PART_PAGE_FIRST>VALUE"`
	ManifestationOfficialJournalPartPageLast  string                            `xml:"MANIFESTATION_OFFICIAL-JOURNAL_PART_PAGE_LAST>VALUE"`
	ManifestationPartOfManifestation          *ManifestationPartOfManifestation `xml:"MANIFESTATION_PART_OF_MANIFESTATION"`
}

type ManifestationPartOfManifestation struct {
	XMLName xml.Name `xml:"MANIFESTATION_PART_OF_MANIFESTATION"`
	Sameas  *Sameas  `xml:"SAMEAS"`
}
