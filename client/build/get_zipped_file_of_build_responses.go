// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetZippedFileOfBuildReader is a Reader for the GetZippedFileOfBuild structure.
type GetZippedFileOfBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZippedFileOfBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetZippedFileOfBuildDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetZippedFileOfBuildDefault creates a GetZippedFileOfBuildDefault with default headers values
func NewGetZippedFileOfBuildDefault(code int) *GetZippedFileOfBuildDefault {
	return &GetZippedFileOfBuildDefault{
		_statusCode: code,
	}
}

/* GetZippedFileOfBuildDefault describes a response with status code -1, with default header values.

successful operation
*/
type GetZippedFileOfBuildDefault struct {
	_statusCode int
}

// Code gets the status code for the get zipped file of build default response
func (o *GetZippedFileOfBuildDefault) Code() int {
	return o._statusCode
}

func (o *GetZippedFileOfBuildDefault) Error() string {
	return fmt.Sprintf("[GET /app/rest/builds/{buildLocator}/artifacts/archived{path}][%d] getZippedFileOfBuild default ", o._statusCode)
}

func (o *GetZippedFileOfBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
