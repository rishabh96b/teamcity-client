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

// AddAgentToAgentPoolReader is a Reader for the AddAgentToAgentPool structure.
type AddAgentToAgentPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAgentToAgentPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddAgentToAgentPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddAgentToAgentPoolOK creates a AddAgentToAgentPoolOK with default headers values
func NewAddAgentToAgentPoolOK() *AddAgentToAgentPoolOK {
	return &AddAgentToAgentPoolOK{}
}

/* AddAgentToAgentPoolOK describes a response with status code 200, with default header values.

successful operation
*/
type AddAgentToAgentPoolOK struct {
	Payload *models.Agent
}

func (o *AddAgentToAgentPoolOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/agentPools/{agentPoolLocator}/agents][%d] addAgentToAgentPoolOK  %+v", 200, o.Payload)
}
func (o *AddAgentToAgentPoolOK) GetPayload() *models.Agent {
	return o.Payload
}

func (o *AddAgentToAgentPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Agent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}