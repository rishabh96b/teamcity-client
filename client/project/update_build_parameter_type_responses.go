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

// UpdateBuildParameterTypeReader is a Reader for the UpdateBuildParameterType structure.
type UpdateBuildParameterTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBuildParameterTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBuildParameterTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateBuildParameterTypeOK creates a UpdateBuildParameterTypeOK with default headers values
func NewUpdateBuildParameterTypeOK() *UpdateBuildParameterTypeOK {
	return &UpdateBuildParameterTypeOK{}
}

/* UpdateBuildParameterTypeOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateBuildParameterTypeOK struct {
	Payload *models.Type
}

func (o *UpdateBuildParameterTypeOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/projects/{projectLocator}/parameters/{name}/type][%d] updateBuildParameterTypeOK  %+v", 200, o.Payload)
}
func (o *UpdateBuildParameterTypeOK) GetPayload() *models.Type {
	return o.Payload
}

func (o *UpdateBuildParameterTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Type)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
