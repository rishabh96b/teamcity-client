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

// UpdateBuildParameterReader is a Reader for the UpdateBuildParameter structure.
type UpdateBuildParameterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBuildParameterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBuildParameterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateBuildParameterOK creates a UpdateBuildParameterOK with default headers values
func NewUpdateBuildParameterOK() *UpdateBuildParameterOK {
	return &UpdateBuildParameterOK{}
}

/* UpdateBuildParameterOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateBuildParameterOK struct {
	Payload *models.Property
}

func (o *UpdateBuildParameterOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/projects/{projectLocator}/parameters/{name}][%d] updateBuildParameterOK  %+v", 200, o.Payload)
}
func (o *UpdateBuildParameterOK) GetPayload() *models.Property {
	return o.Payload
}

func (o *UpdateBuildParameterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Property)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}