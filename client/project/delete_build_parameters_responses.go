// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteBuildParametersReader is a Reader for the DeleteBuildParameters structure.
type DeleteBuildParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBuildParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteBuildParametersDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteBuildParametersDefault creates a DeleteBuildParametersDefault with default headers values
func NewDeleteBuildParametersDefault(code int) *DeleteBuildParametersDefault {
	return &DeleteBuildParametersDefault{
		_statusCode: code,
	}
}

/* DeleteBuildParametersDefault describes a response with status code -1, with default header values.

successful operation
*/
type DeleteBuildParametersDefault struct {
	_statusCode int
}

// Code gets the status code for the delete build parameters default response
func (o *DeleteBuildParametersDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBuildParametersDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/projects/{projectLocator}/parameters][%d] deleteBuildParameters default ", o._statusCode)
}

func (o *DeleteBuildParametersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
