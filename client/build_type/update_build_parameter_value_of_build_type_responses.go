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

// UpdateBuildParameterValueOfBuildTypeReader is a Reader for the UpdateBuildParameterValueOfBuildType structure.
type UpdateBuildParameterValueOfBuildTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBuildParameterValueOfBuildTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBuildParameterValueOfBuildTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateBuildParameterValueOfBuildTypeOK creates a UpdateBuildParameterValueOfBuildTypeOK with default headers values
func NewUpdateBuildParameterValueOfBuildTypeOK() *UpdateBuildParameterValueOfBuildTypeOK {
	return &UpdateBuildParameterValueOfBuildTypeOK{}
}

/* UpdateBuildParameterValueOfBuildTypeOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateBuildParameterValueOfBuildTypeOK struct {
	Payload string
}

func (o *UpdateBuildParameterValueOfBuildTypeOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/parameters/{name}/value][%d] updateBuildParameterValueOfBuildTypeOK  %+v", 200, o.Payload)
}
func (o *UpdateBuildParameterValueOfBuildTypeOK) GetPayload() string {
	return o.Payload
}

func (o *UpdateBuildParameterValueOfBuildTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
