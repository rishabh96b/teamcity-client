// Code generated by go-swagger; DO NOT EDIT.

package vcs_root_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteVcsRootInstanceRepositoryStateReader is a Reader for the DeleteVcsRootInstanceRepositoryState structure.
type DeleteVcsRootInstanceRepositoryStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVcsRootInstanceRepositoryStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteVcsRootInstanceRepositoryStateDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteVcsRootInstanceRepositoryStateDefault creates a DeleteVcsRootInstanceRepositoryStateDefault with default headers values
func NewDeleteVcsRootInstanceRepositoryStateDefault(code int) *DeleteVcsRootInstanceRepositoryStateDefault {
	return &DeleteVcsRootInstanceRepositoryStateDefault{
		_statusCode: code,
	}
}

/* DeleteVcsRootInstanceRepositoryStateDefault describes a response with status code -1, with default header values.

successful operation
*/
type DeleteVcsRootInstanceRepositoryStateDefault struct {
	_statusCode int
}

// Code gets the status code for the delete vcs root instance repository state default response
func (o *DeleteVcsRootInstanceRepositoryStateDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVcsRootInstanceRepositoryStateDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/repositoryState][%d] deleteVcsRootInstanceRepositoryState default ", o._statusCode)
}

func (o *DeleteVcsRootInstanceRepositoryStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
