// Code generated by go-swagger; DO NOT EDIT.

package root

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRootEndpointsOfRootReader is a Reader for the GetRootEndpointsOfRoot structure.
type GetRootEndpointsOfRootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRootEndpointsOfRootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRootEndpointsOfRootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRootEndpointsOfRootOK creates a GetRootEndpointsOfRootOK with default headers values
func NewGetRootEndpointsOfRootOK() *GetRootEndpointsOfRootOK {
	return &GetRootEndpointsOfRootOK{}
}

/* GetRootEndpointsOfRootOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRootEndpointsOfRootOK struct {
	Payload string
}

func (o *GetRootEndpointsOfRootOK) Error() string {
	return fmt.Sprintf("[GET /app/rest][%d] getRootEndpointsOfRootOK  %+v", 200, o.Payload)
}
func (o *GetRootEndpointsOfRootOK) GetPayload() string {
	return o.Payload
}

func (o *GetRootEndpointsOfRootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
