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

// SetBuildNumberReader is a Reader for the SetBuildNumber structure.
type SetBuildNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetBuildNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetBuildNumberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetBuildNumberOK creates a SetBuildNumberOK with default headers values
func NewSetBuildNumberOK() *SetBuildNumberOK {
	return &SetBuildNumberOK{}
}

/* SetBuildNumberOK describes a response with status code 200, with default header values.

successful operation
*/
type SetBuildNumberOK struct {
	Payload string
}

func (o *SetBuildNumberOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/builds/{buildLocator}/number][%d] setBuildNumberOK  %+v", 200, o.Payload)
}
func (o *SetBuildNumberOK) GetPayload() string {
	return o.Payload
}

func (o *SetBuildNumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
