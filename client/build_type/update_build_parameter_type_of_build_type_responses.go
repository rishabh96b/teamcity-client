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

// UpdateBuildParameterTypeOfBuildTypeReader is a Reader for the UpdateBuildParameterTypeOfBuildType structure.
type UpdateBuildParameterTypeOfBuildTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBuildParameterTypeOfBuildTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBuildParameterTypeOfBuildTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateBuildParameterTypeOfBuildTypeOK creates a UpdateBuildParameterTypeOfBuildTypeOK with default headers values
func NewUpdateBuildParameterTypeOfBuildTypeOK() *UpdateBuildParameterTypeOfBuildTypeOK {
	return &UpdateBuildParameterTypeOfBuildTypeOK{}
}

/* UpdateBuildParameterTypeOfBuildTypeOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateBuildParameterTypeOfBuildTypeOK struct {
	Payload *models.Type
}

func (o *UpdateBuildParameterTypeOfBuildTypeOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/parameters/{name}/type][%d] updateBuildParameterTypeOfBuildTypeOK  %+v", 200, o.Payload)
}
func (o *UpdateBuildParameterTypeOfBuildTypeOK) GetPayload() *models.Type {
	return o.Payload
}

func (o *UpdateBuildParameterTypeOfBuildTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Type)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}