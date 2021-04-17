package webservice

import (
	"fmt"

	"github.com/slipke/eurlex-ws-go/pkg/response"
)

type ErrorResponse struct {
	Code   string
	Reason string
}

func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}

func NewErrorResponseFromXML(xml string) (*ErrorResponse, error) {
	er := NewErrorResponse()
	e, err := response.NewEnvelopeFromXML(xml)
	if err != nil {
		return nil, fmt.Errorf("failed to create envelope from XML: %s", err)
	}

	er.Code = e.Body.Fault.Code
	er.Reason = e.Body.Fault.Reason

	return er, nil
}
