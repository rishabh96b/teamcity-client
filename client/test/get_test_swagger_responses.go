// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// GetTestReader is a Reader for the GetTest structure.
type GetTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTestOK creates a GetTestOK with default headers values
func NewGetTestOK() *GetTestOK {
	return &GetTestOK{}
}

/* GetTestOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTestOK struct {
	Payload *models.Test
}

func (o *GetTestOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/tests/{testLocator}][%d] getTestOK  %+v", 200, o.Payload)
}
func (o *GetTestOK) GetPayload() *models.Test {
	return o.Payload
}

func (o *GetTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Test)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}