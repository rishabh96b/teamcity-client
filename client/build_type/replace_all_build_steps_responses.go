// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// ReplaceAllBuildStepsReader is a Reader for the ReplaceAllBuildSteps structure.
type ReplaceAllBuildStepsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceAllBuildStepsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceAllBuildStepsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceAllBuildStepsOK creates a ReplaceAllBuildStepsOK with default headers values
func NewReplaceAllBuildStepsOK() *ReplaceAllBuildStepsOK {
	return &ReplaceAllBuildStepsOK{}
}

/* ReplaceAllBuildStepsOK describes a response with status code 200, with default header values.

successful operation
*/
type ReplaceAllBuildStepsOK struct {
	Payload *models.Steps
}

func (o *ReplaceAllBuildStepsOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/steps][%d] replaceAllBuildStepsOK  %+v", 200, o.Payload)
}
func (o *ReplaceAllBuildStepsOK) GetPayload() *models.Steps {
	return o.Payload
}

func (o *ReplaceAllBuildStepsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Steps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
