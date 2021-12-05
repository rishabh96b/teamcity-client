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

// AddProjectToAgentPoolReader is a Reader for the AddProjectToAgentPool structure.
type AddProjectToAgentPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProjectToAgentPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddProjectToAgentPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddProjectToAgentPoolOK creates a AddProjectToAgentPoolOK with default headers values
func NewAddProjectToAgentPoolOK() *AddProjectToAgentPoolOK {
	return &AddProjectToAgentPoolOK{}
}

/* AddProjectToAgentPoolOK describes a response with status code 200, with default header values.

successful operation
*/
type AddProjectToAgentPoolOK struct {
	Payload *models.Project
}

func (o *AddProjectToAgentPoolOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/agentPools/{agentPoolLocator}/projects][%d] addProjectToAgentPoolOK  %+v", 200, o.Payload)
}
func (o *AddProjectToAgentPoolOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *AddProjectToAgentPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}