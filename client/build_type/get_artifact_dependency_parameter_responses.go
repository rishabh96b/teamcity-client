// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetArtifactDependencyParameterReader is a Reader for the GetArtifactDependencyParameter structure.
type GetArtifactDependencyParameterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArtifactDependencyParameterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArtifactDependencyParameterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArtifactDependencyParameterOK creates a GetArtifactDependencyParameterOK with default headers values
func NewGetArtifactDependencyParameterOK() *GetArtifactDependencyParameterOK {
	return &GetArtifactDependencyParameterOK{}
}

/* GetArtifactDependencyParameterOK describes a response with status code 200, with default header values.

successful operation
*/
type GetArtifactDependencyParameterOK struct {
	Payload string
}

func (o *GetArtifactDependencyParameterOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/artifact-dependencies/{artifactDepLocator}/{fieldName}][%d] getArtifactDependencyParameterOK  %+v", 200, o.Payload)
}
func (o *GetArtifactDependencyParameterOK) GetPayload() string {
	return o.Payload
}

func (o *GetArtifactDependencyParameterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
