// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetBuildStatusTextReader is a Reader for the GetBuildStatusText structure.
type GetBuildStatusTextReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildStatusTextReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildStatusTextOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBuildStatusTextOK creates a GetBuildStatusTextOK with default headers values
func NewGetBuildStatusTextOK() *GetBuildStatusTextOK {
	return &GetBuildStatusTextOK{}
}

/* GetBuildStatusTextOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBuildStatusTextOK struct {
	Payload string
}

func (o *GetBuildStatusTextOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/builds/{buildLocator}/statusText][%d] getBuildStatusTextOK  %+v", 200, o.Payload)
}
func (o *GetBuildStatusTextOK) GetPayload() string {
	return o.Payload
}

func (o *GetBuildStatusTextOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}