// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSnapshotDependencyReader is a Reader for the DeleteSnapshotDependency structure.
type DeleteSnapshotDependencyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSnapshotDependencyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteSnapshotDependencyDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteSnapshotDependencyDefault creates a DeleteSnapshotDependencyDefault with default headers values
func NewDeleteSnapshotDependencyDefault(code int) *DeleteSnapshotDependencyDefault {
	return &DeleteSnapshotDependencyDefault{
		_statusCode: code,
	}
}

/* DeleteSnapshotDependencyDefault describes a response with status code -1, with default header values.

successful operation
*/
type DeleteSnapshotDependencyDefault struct {
	_statusCode int
}

// Code gets the status code for the delete snapshot dependency default response
func (o *DeleteSnapshotDependencyDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSnapshotDependencyDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/buildTypes/{btLocator}/snapshot-dependencies/{snapshotDepLocator}][%d] deleteSnapshotDependency default ", o._statusCode)
}

func (o *DeleteSnapshotDependencyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
