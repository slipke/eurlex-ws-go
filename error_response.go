package eurlex

import (
	"fmt"

	"github.com/slipke/eurlex-ws-go/internal/response"
)

const (
	errorFailedToCreateResponse = "failed to create envelope from XML"
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
		return nil, fmt.Errorf("%s: %s", errorFailedToCreateResponse, err)
	}

	er.Code = e.Body.Fault.Code
	er.Reason = e.Body.Fault.Reason

	return er, nil
}
