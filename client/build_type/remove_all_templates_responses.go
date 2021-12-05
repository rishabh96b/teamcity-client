// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveAllTemplatesReader is a Reader for the RemoveAllTemplates structure.
type RemoveAllTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveAllTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewRemoveAllTemplatesDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewRemoveAllTemplatesDefault creates a RemoveAllTemplatesDefault with default headers values
func NewRemoveAllTemplatesDefault(code int) *RemoveAllTemplatesDefault {
	return &RemoveAllTemplatesDefault{
		_statusCode: code,
	}
}

/* RemoveAllTemplatesDefault describes a response with status code -1, with default header values.

successful operation
*/
type RemoveAllTemplatesDefault struct {
	_statusCode int
}

// Code gets the status code for the remove all templates default response
func (o *RemoveAllTemplatesDefault) Code() int {
	return o._statusCode
}

func (o *RemoveAllTemplatesDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/buildTypes/{btLocator}/templates][%d] removeAllTemplates default ", o._statusCode)
}

func (o *RemoveAllTemplatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
