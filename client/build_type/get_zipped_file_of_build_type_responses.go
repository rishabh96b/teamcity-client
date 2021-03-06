// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetZippedFileOfBuildTypeReader is a Reader for the GetZippedFileOfBuildType structure.
type GetZippedFileOfBuildTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZippedFileOfBuildTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetZippedFileOfBuildTypeDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetZippedFileOfBuildTypeDefault creates a GetZippedFileOfBuildTypeDefault with default headers values
func NewGetZippedFileOfBuildTypeDefault(code int) *GetZippedFileOfBuildTypeDefault {
	return &GetZippedFileOfBuildTypeDefault{
		_statusCode: code,
	}
}

/* GetZippedFileOfBuildTypeDefault describes a response with status code -1, with default header values.

successful operation
*/
type GetZippedFileOfBuildTypeDefault struct {
	_statusCode int
}

// Code gets the status code for the get zipped file of build type default response
func (o *GetZippedFileOfBuildTypeDefault) Code() int {
	return o._statusCode
}

func (o *GetZippedFileOfBuildTypeDefault) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/vcs/files/latest/archived{path}][%d] getZippedFileOfBuildType default ", o._statusCode)
}

func (o *GetZippedFileOfBuildTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
