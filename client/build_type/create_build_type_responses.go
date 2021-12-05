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

// CreateBuildTypeReader is a Reader for the CreateBuildType structure.
type CreateBuildTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBuildTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBuildTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateBuildTypeOK creates a CreateBuildTypeOK with default headers values
func NewCreateBuildTypeOK() *CreateBuildTypeOK {
	return &CreateBuildTypeOK{}
}

/* CreateBuildTypeOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateBuildTypeOK struct {
	Payload *models.BuildType
}

func (o *CreateBuildTypeOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/buildTypes][%d] createBuildTypeOK  %+v", 200, o.Payload)
}
func (o *CreateBuildTypeOK) GetPayload() *models.BuildType {
	return o.Payload
}

func (o *CreateBuildTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}