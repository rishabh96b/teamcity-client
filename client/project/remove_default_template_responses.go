// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveDefaultTemplateReader is a Reader for the RemoveDefaultTemplate structure.
type RemoveDefaultTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveDefaultTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewRemoveDefaultTemplateDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewRemoveDefaultTemplateDefault creates a RemoveDefaultTemplateDefault with default headers values
func NewRemoveDefaultTemplateDefault(code int) *RemoveDefaultTemplateDefault {
	return &RemoveDefaultTemplateDefault{
		_statusCode: code,
	}
}

/* RemoveDefaultTemplateDefault describes a response with status code -1, with default header values.

successful operation
*/
type RemoveDefaultTemplateDefault struct {
	_statusCode int
}

// Code gets the status code for the remove default template default response
func (o *RemoveDefaultTemplateDefault) Code() int {
	return o._statusCode
}

func (o *RemoveDefaultTemplateDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/projects/{projectLocator}/defaultTemplate][%d] removeDefaultTemplate default ", o._statusCode)
}

func (o *RemoveDefaultTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
