// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProjectFromAgentPoolReader is a Reader for the DeleteProjectFromAgentPool structure.
type DeleteProjectFromAgentPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectFromAgentPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteProjectFromAgentPoolDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteProjectFromAgentPoolDefault creates a DeleteProjectFromAgentPoolDefault with default headers values
func NewDeleteProjectFromAgentPoolDefault(code int) *DeleteProjectFromAgentPoolDefault {
	return &DeleteProjectFromAgentPoolDefault{
		_statusCode: code,
	}
}

/* DeleteProjectFromAgentPoolDefault describes a response with status code -1, with default header values.

successful operation
*/
type DeleteProjectFromAgentPoolDefault struct {
	_statusCode int
}

// Code gets the status code for the delete project from agent pool default response
func (o *DeleteProjectFromAgentPoolDefault) Code() int {
	return o._statusCode
}

func (o *DeleteProjectFromAgentPoolDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/agentPools/{agentPoolLocator}/projects/{projectLocator}][%d] deleteProjectFromAgentPool default ", o._statusCode)
}

func (o *DeleteProjectFromAgentPoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}