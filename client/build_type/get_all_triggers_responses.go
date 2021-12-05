// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// GetAllTriggersReader is a Reader for the GetAllTriggers structure.
type GetAllTriggersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTriggersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllTriggersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllTriggersOK creates a GetAllTriggersOK with default headers values
func NewGetAllTriggersOK() *GetAllTriggersOK {
	return &GetAllTriggersOK{}
}

/* GetAllTriggersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllTriggersOK struct {
	Payload *models.Triggers
}

func (o *GetAllTriggersOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/triggers][%d] getAllTriggersOK  %+v", 200, o.Payload)
}
func (o *GetAllTriggersOK) GetPayload() *models.Triggers {
	return o.Payload
}

func (o *GetAllTriggersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Triggers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}