// Code generated by go-swagger; DO NOT EDIT.

package investigation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// GetInvestigationReader is a Reader for the GetInvestigation structure.
type GetInvestigationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvestigationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInvestigationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInvestigationOK creates a GetInvestigationOK with default headers values
func NewGetInvestigationOK() *GetInvestigationOK {
	return &GetInvestigationOK{}
}

/* GetInvestigationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvestigationOK struct {
	Payload *models.Investigation
}

func (o *GetInvestigationOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/investigations/{investigationLocator}][%d] getInvestigationOK  %+v", 200, o.Payload)
}
func (o *GetInvestigationOK) GetPayload() *models.Investigation {
	return o.Payload
}

func (o *GetInvestigationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Investigation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}