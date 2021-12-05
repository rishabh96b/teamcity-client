// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveUserPropertyReader is a Reader for the RemoveUserProperty structure.
type RemoveUserPropertyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveUserPropertyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewRemoveUserPropertyDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewRemoveUserPropertyDefault creates a RemoveUserPropertyDefault with default headers values
func NewRemoveUserPropertyDefault(code int) *RemoveUserPropertyDefault {
	return &RemoveUserPropertyDefault{
		_statusCode: code,
	}
}

/* RemoveUserPropertyDefault describes a response with status code -1, with default header values.

successful operation
*/
type RemoveUserPropertyDefault struct {
	_statusCode int
}

// Code gets the status code for the remove user property default response
func (o *RemoveUserPropertyDefault) Code() int {
	return o._statusCode
}

func (o *RemoveUserPropertyDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/users/{userLocator}/properties/{name}][%d] removeUserProperty default ", o._statusCode)
}

func (o *RemoveUserPropertyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
