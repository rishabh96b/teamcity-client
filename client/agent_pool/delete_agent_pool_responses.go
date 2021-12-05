// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAgentPoolReader is a Reader for the DeleteAgentPool structure.
type DeleteAgentPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAgentPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteAgentPoolDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteAgentPoolDefault creates a DeleteAgentPoolDefault with default headers values
func NewDeleteAgentPoolDefault(code int) *DeleteAgentPoolDefault {
	return &DeleteAgentPoolDefault{
		_statusCode: code,
	}
}

/* DeleteAgentPoolDefault describes a response with status code -1, with default header values.

successful operation
*/
type DeleteAgentPoolDefault struct {
	_statusCode int
}

// Code gets the status code for the delete agent pool default response
func (o *DeleteAgentPoolDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAgentPoolDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/agentPools/{agentPoolLocator}][%d] deleteAgentPool default ", o._statusCode)
}

func (o *DeleteAgentPoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
