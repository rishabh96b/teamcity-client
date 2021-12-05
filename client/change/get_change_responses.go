// Code generated by go-swagger; DO NOT EDIT.

package change

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// GetChangeReader is a Reader for the GetChange structure.
type GetChangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetChangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetChangeOK creates a GetChangeOK with default headers values
func NewGetChangeOK() *GetChangeOK {
	return &GetChangeOK{}
}

/* GetChangeOK describes a response with status code 200, with default header values.

successful operation
*/
type GetChangeOK struct {
	Payload *models.Change
}

func (o *GetChangeOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/changes/{changeLocator}][%d] getChangeOK  %+v", 200, o.Payload)
}
func (o *GetChangeOK) GetPayload() *models.Change {
	return o.Payload
}

func (o *GetChangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Change)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}