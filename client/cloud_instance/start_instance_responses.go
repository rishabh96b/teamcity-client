// Code generated by go-swagger; DO NOT EDIT.

package cloud_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StartInstanceReader is a Reader for the StartInstance structure.
type StartInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewStartInstanceDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewStartInstanceDefault creates a StartInstanceDefault with default headers values
func NewStartInstanceDefault(code int) *StartInstanceDefault {
	return &StartInstanceDefault{
		_statusCode: code,
	}
}

/* StartInstanceDefault describes a response with status code -1, with default header values.

successful operation
*/
type StartInstanceDefault struct {
	_statusCode int
}

// Code gets the status code for the start instance default response
func (o *StartInstanceDefault) Code() int {
	return o._statusCode
}

func (o *StartInstanceDefault) Error() string {
	return fmt.Sprintf("[POST /app/rest/cloud/instances][%d] startInstance default ", o._statusCode)
}

func (o *StartInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
