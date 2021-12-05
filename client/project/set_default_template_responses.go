// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// SetDefaultTemplateReader is a Reader for the SetDefaultTemplate structure.
type SetDefaultTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDefaultTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetDefaultTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetDefaultTemplateOK creates a SetDefaultTemplateOK with default headers values
func NewSetDefaultTemplateOK() *SetDefaultTemplateOK {
	return &SetDefaultTemplateOK{}
}

/* SetDefaultTemplateOK describes a response with status code 200, with default header values.

successful operation
*/
type SetDefaultTemplateOK struct {
	Payload *models.BuildType
}

func (o *SetDefaultTemplateOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/projects/{projectLocator}/defaultTemplate][%d] setDefaultTemplateOK  %+v", 200, o.Payload)
}
func (o *SetDefaultTemplateOK) GetPayload() *models.BuildType {
	return o.Payload
}

func (o *SetDefaultTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
