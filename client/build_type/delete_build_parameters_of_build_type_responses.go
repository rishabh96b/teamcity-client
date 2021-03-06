// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteBuildParametersOfBuildTypeReader is a Reader for the DeleteBuildParametersOfBuildType structure.
type DeleteBuildParametersOfBuildTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBuildParametersOfBuildTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteBuildParametersOfBuildTypeDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteBuildParametersOfBuildTypeDefault creates a DeleteBuildParametersOfBuildTypeDefault with default headers values
func NewDeleteBuildParametersOfBuildTypeDefault(code int) *DeleteBuildParametersOfBuildTypeDefault {
	return &DeleteBuildParametersOfBuildTypeDefault{
		_statusCode: code,
	}
}

/* DeleteBuildParametersOfBuildTypeDefault describes a response with status code -1, with default header values.

successful operation
*/
type DeleteBuildParametersOfBuildTypeDefault struct {
	_statusCode int
}

// Code gets the status code for the delete build parameters of build type default response
func (o *DeleteBuildParametersOfBuildTypeDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBuildParametersOfBuildTypeDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/buildTypes/{btLocator}/parameters][%d] deleteBuildParametersOfBuildType default ", o._statusCode)
}

func (o *DeleteBuildParametersOfBuildTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
