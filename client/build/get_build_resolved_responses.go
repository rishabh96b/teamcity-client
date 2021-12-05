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

// GetBuildResolvedReader is a Reader for the GetBuildResolved structure.
type GetBuildResolvedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildResolvedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildResolvedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBuildResolvedOK creates a GetBuildResolvedOK with default headers values
func NewGetBuildResolvedOK() *GetBuildResolvedOK {
	return &GetBuildResolvedOK{}
}

/* GetBuildResolvedOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBuildResolvedOK struct {
	Payload string
}

func (o *GetBuildResolvedOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/builds/{buildLocator}/resolved/{value}][%d] getBuildResolvedOK  %+v", 200, o.Payload)
}
func (o *GetBuildResolvedOK) GetPayload() string {
	return o.Payload
}

func (o *GetBuildResolvedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}