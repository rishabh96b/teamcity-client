// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetGroupPropertyReader is a Reader for the GetGroupProperty structure.
type GetGroupPropertyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupPropertyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupPropertyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupPropertyOK creates a GetGroupPropertyOK with default headers values
func NewGetGroupPropertyOK() *GetGroupPropertyOK {
	return &GetGroupPropertyOK{}
}

/* GetGroupPropertyOK describes a response with status code 200, with default header values.

successful operation
*/
type GetGroupPropertyOK struct {
	Payload string
}

func (o *GetGroupPropertyOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/userGroups/{groupLocator}/properties/{name}][%d] getGroupPropertyOK  %+v", 200, o.Payload)
}
func (o *GetGroupPropertyOK) GetPayload() string {
	return o.Payload
}

func (o *GetGroupPropertyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
