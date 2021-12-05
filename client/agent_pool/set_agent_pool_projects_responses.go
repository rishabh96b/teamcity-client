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

// SetAgentPoolProjectsReader is a Reader for the SetAgentPoolProjects structure.
type SetAgentPoolProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAgentPoolProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetAgentPoolProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetAgentPoolProjectsOK creates a SetAgentPoolProjectsOK with default headers values
func NewSetAgentPoolProjectsOK() *SetAgentPoolProjectsOK {
	return &SetAgentPoolProjectsOK{}
}

/* SetAgentPoolProjectsOK describes a response with status code 200, with default header values.

successful operation
*/
type SetAgentPoolProjectsOK struct {
	Payload *models.Projects
}

func (o *SetAgentPoolProjectsOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/agentPools/{agentPoolLocator}/projects][%d] setAgentPoolProjectsOK  %+v", 200, o.Payload)
}
func (o *SetAgentPoolProjectsOK) GetPayload() *models.Projects {
	return o.Payload
}

func (o *SetAgentPoolProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Projects)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
