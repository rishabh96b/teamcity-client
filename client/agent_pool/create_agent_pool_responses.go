// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// CreateAgentPoolReader is a Reader for the CreateAgentPool structure.
type CreateAgentPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAgentPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAgentPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAgentPoolOK creates a CreateAgentPoolOK with default headers values
func NewCreateAgentPoolOK() *CreateAgentPoolOK {
	return &CreateAgentPoolOK{}
}

/* CreateAgentPoolOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateAgentPoolOK struct {
	Payload *models.AgentPool
}

func (o *CreateAgentPoolOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/agentPools][%d] createAgentPoolOK  %+v", 200, o.Payload)
}
func (o *CreateAgentPoolOK) GetPayload() *models.AgentPool {
	return o.Payload
}

func (o *CreateAgentPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AgentPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
