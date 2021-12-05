// Code generated by go-swagger; DO NOT EDIT.

package agent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// GetIncompatibleBuildTypesReader is a Reader for the GetIncompatibleBuildTypes structure.
type GetIncompatibleBuildTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIncompatibleBuildTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIncompatibleBuildTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIncompatibleBuildTypesOK creates a GetIncompatibleBuildTypesOK with default headers values
func NewGetIncompatibleBuildTypesOK() *GetIncompatibleBuildTypesOK {
	return &GetIncompatibleBuildTypesOK{}
}

/* GetIncompatibleBuildTypesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetIncompatibleBuildTypesOK struct {
	Payload *models.Compatibilities
}

func (o *GetIncompatibleBuildTypesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/agents/{agentLocator}/incompatibleBuildTypes][%d] getIncompatibleBuildTypesOK  %+v", 200, o.Payload)
}
func (o *GetIncompatibleBuildTypesOK) GetPayload() *models.Compatibilities {
	return o.Payload
}

func (o *GetIncompatibleBuildTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Compatibilities)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
