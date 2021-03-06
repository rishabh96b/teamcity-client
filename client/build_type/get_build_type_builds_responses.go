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

// GetBuildTypeBuildsReader is a Reader for the GetBuildTypeBuilds structure.
type GetBuildTypeBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildTypeBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildTypeBuildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBuildTypeBuildsOK creates a GetBuildTypeBuildsOK with default headers values
func NewGetBuildTypeBuildsOK() *GetBuildTypeBuildsOK {
	return &GetBuildTypeBuildsOK{}
}

/* GetBuildTypeBuildsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBuildTypeBuildsOK struct {
	Payload *models.Builds
}

func (o *GetBuildTypeBuildsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/builds][%d] getBuildTypeBuildsOK  %+v", 200, o.Payload)
}
func (o *GetBuildTypeBuildsOK) GetPayload() *models.Builds {
	return o.Payload
}

func (o *GetBuildTypeBuildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Builds)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
