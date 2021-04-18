package webservice

import "testing"

const (
	errorRes = `<?xml version='1.0' encoding='UTF-8'?>
    <S:Envelope xmlns:S="http://www.w3.org/2003/05/soap-envelope">
        <S:Header>
            <NotUnderstood xmlns:abc="http://docs.oasisopen.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
                xmlns="http://www.w3.org/2003/05/soap-envelope" qname="abc:Security"/>
        </S:Header>
        <S:Body>
            <ns1:Fault xmlns:ns0="http://schemas.xmlsoap.org/soap/envelope/"
                xmlns:ns1="http://www.w3.org/2003/05/soap-envelope">
                <ns1:Code>
                    <ns1:Value>ns1:MustUnderstand</ns1:Value>
                </ns1:Code>
                <ns1:Reason>
                    <ns1:Text xml:lang="en">One or more mandatory SOAP header blocks not understood</ns1:Text>
                </ns1:Reason>
            </ns1:Fault>
        </S:Body>
    </S:Envelope>`
)

func TestErrorResponseFromXML(t *testing.T) {
	er, err := NewErrorResponseFromXML(errorRes)
	if err != nil {
		t.Errorf("Failed to unmarshal error response from XML: %s", err)
		t.FailNow()
	}

	if er.Code != "ns1:MustUnderstand" {
		t.Errorf("Invalid field Code, got %s, want %s", er.Code, "ns1:MustUnderstand")
	}

	if er.Reason != "One or more mandatory SOAP header blocks not understood" {
		t.Errorf("Invalid field Reason, got %s, want: %s", er.Reason, "One or more mandatory SOAP header blocks not understood")
	}

}
