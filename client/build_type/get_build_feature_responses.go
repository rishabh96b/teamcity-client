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

// GetBuildFeatureReader is a Reader for the GetBuildFeature structure.
type GetBuildFeatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildFeatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildFeatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBuildFeatureOK creates a GetBuildFeatureOK with default headers values
func NewGetBuildFeatureOK() *GetBuildFeatureOK {
	return &GetBuildFeatureOK{}
}

/* GetBuildFeatureOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBuildFeatureOK struct {
	Payload *models.Feature
}

func (o *GetBuildFeatureOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/features/{featureId}][%d] getBuildFeatureOK  %+v", 200, o.Payload)
}
func (o *GetBuildFeatureOK) GetPayload() *models.Feature {
	return o.Payload
}

func (o *GetBuildFeatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Feature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
