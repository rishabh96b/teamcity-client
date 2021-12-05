// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetBuildStatusIconReader is a Reader for the GetBuildStatusIcon structure.
type GetBuildStatusIconReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildStatusIconReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetBuildStatusIconDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetBuildStatusIconDefault creates a GetBuildStatusIconDefault with default headers values
func NewGetBuildStatusIconDefault(code int) *GetBuildStatusIconDefault {
	return &GetBuildStatusIconDefault{
		_statusCode: code,
	}
}

/* GetBuildStatusIconDefault describes a response with status code -1, with default header values.

successful operation
*/
type GetBuildStatusIconDefault struct {
	_statusCode int
}

// Code gets the status code for the get build status icon default response
func (o *GetBuildStatusIconDefault) Code() int {
	return o._statusCode
}

func (o *GetBuildStatusIconDefault) Error() string {
	return fmt.Sprintf("[GET /app/rest/builds/{buildLocator}/statusIcon{suffix}][%d] getBuildStatusIcon default ", o._statusCode)
}

func (o *GetBuildStatusIconDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}