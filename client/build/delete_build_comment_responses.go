// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteBuildCommentReader is a Reader for the DeleteBuildComment structure.
type DeleteBuildCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBuildCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteBuildCommentDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteBuildCommentDefault creates a DeleteBuildCommentDefault with default headers values
func NewDeleteBuildCommentDefault(code int) *DeleteBuildCommentDefault {
	return &DeleteBuildCommentDefault{
		_statusCode: code,
	}
}

/* DeleteBuildCommentDefault describes a response with status code -1, with default header values.

successful operation
*/
type DeleteBuildCommentDefault struct {
	_statusCode int
}

// Code gets the status code for the delete build comment default response
func (o *DeleteBuildCommentDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBuildCommentDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/builds/{buildLocator}/comment][%d] deleteBuildComment default ", o._statusCode)
}

func (o *DeleteBuildCommentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
