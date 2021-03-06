// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteVcsRootOfBuildTypeReader is a Reader for the DeleteVcsRootOfBuildType structure.
type DeleteVcsRootOfBuildTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVcsRootOfBuildTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteVcsRootOfBuildTypeDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteVcsRootOfBuildTypeDefault creates a DeleteVcsRootOfBuildTypeDefault with default headers values
func NewDeleteVcsRootOfBuildTypeDefault(code int) *DeleteVcsRootOfBuildTypeDefault {
	return &DeleteVcsRootOfBuildTypeDefault{
		_statusCode: code,
	}
}

/* DeleteVcsRootOfBuildTypeDefault describes a response with status code -1, with default header values.

successful operation
*/
type DeleteVcsRootOfBuildTypeDefault struct {
	_statusCode int
}

// Code gets the status code for the delete vcs root of build type default response
func (o *DeleteVcsRootOfBuildTypeDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVcsRootOfBuildTypeDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/buildTypes/{btLocator}/vcs-root-entries/{vcsRootLocator}][%d] deleteVcsRootOfBuildType default ", o._statusCode)
}

func (o *DeleteVcsRootOfBuildTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
