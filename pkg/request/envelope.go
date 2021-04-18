package request

import (
	"encoding/xml"
)

type Envelope struct {
	XMLName          xml.Name `xml:"soap:Envelope"`
	XMLNamespaceSear string   `xml:"xmlns:sear,attr"`
	XMLNAmespaceSoap string   `xml:"xmlns:soap,attr"`
	Header           *Header  `xml:"soap:Header"`
	Body             *Body    `xml:"soap:Body"`
}

func NewEnvelope() *Envelope {
	return &Envelope{
		XMLNamespaceSear: "http://eur-lex.europa.eu/search",
		XMLNAmespaceSoap: "http://www.w3.org/2003/05/soap-envelope",
		Header:           NewHeader(),
		Body:             NewBody(),
	}
}

type Header struct {
	XMLName     xml.Name `xml:"soap:Header"`
	RootElement interface{}
}

func NewHeader() *Header {
	return &Header{}
}

type Security struct {
	XMLName          xml.Name       `xml:"wsse:Security"`
	XMLNamespaceWSSE string         `xml:"xmlns:wsse,attr"`
	MustUnderstand   bool           `xml:"soap:mustUnderstand,attr"`
	UsernameToken    *UsernameToken `xml:"wsse:UsernameToken"`
}

func NewSecurity(u, p string) *Security {
	return &Security{
		XMLNamespaceWSSE: "http://docs.oasisopen.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
		MustUnderstand:   true,
		UsernameToken:    NewUsernameToken(u, p),
	}
}

type UsernameToken struct {
	XMLName         xml.Name  `xml:"wsse:UsernameToken"`
	WSUID           string    `xml:"wsu:Id,attr"`
	XMLNamespaceWSU string    `xml:"xmlns:wsu,attr"`
	Username        string    `xml:"wsse:Username"`
	Password        *Password `xml:"wsse:Password"`
}

func NewUsernameToken(u, p string) *UsernameToken {
	return &UsernameToken{
		WSUID:           "UsernameToken-3",
		XMLNamespaceWSU: "http://docs.oasisopen.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
		Username:        u,
		Password:        NewPassword(p),
	}
}

type Password struct {
	XMLName  xml.Name `xml:"wsse:Password"`
	Type     string   `xml:"Type,attr"`
	Password string   `xml:",chardata"`
}

func NewPassword(p string) *Password {
	return &Password{
		Password: p,
		Type:     "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText",
	}
}

type Body struct {
	XMLName     xml.Name `xml:"soap:Body"`
	RootElement interface{}
}

func NewBody() *Body {
	return &Body{}
}

type SearchRequest struct {
	XMLName        xml.Name `xml:"sear:searchRequest"`
	ExpertQuery    string   `xml:"sear:expertQuery"`
	Page           int64    `xml:"sear:page"`
	PageSize       int64    `xml:"sear:pageSize"`
	SearchLanguage string   `xml:"sear:searchLanguage"`
}

func NewSearchRequest(q string, p, ps int64, lang string) *SearchRequest {
	return &SearchRequest{
		ExpertQuery:    q,
		Page:           p,
		PageSize:       ps,
		SearchLanguage: lang,
	}
}

func (e *Envelope) ToXML() ([]byte, error) {
	xmlBytes, err := xml.Marshal(e)
	if err != nil {
		return nil, err
	}
	return xmlBytes, nil
}
