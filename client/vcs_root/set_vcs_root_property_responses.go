// Code generated by go-swagger; DO NOT EDIT.

package vcs_root

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SetVcsRootPropertyReader is a Reader for the SetVcsRootProperty structure.
type SetVcsRootPropertyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetVcsRootPropertyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetVcsRootPropertyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetVcsRootPropertyOK creates a SetVcsRootPropertyOK with default headers values
func NewSetVcsRootPropertyOK() *SetVcsRootPropertyOK {
	return &SetVcsRootPropertyOK{}
}

/* SetVcsRootPropertyOK describes a response with status code 200, with default header values.

successful operation
*/
type SetVcsRootPropertyOK struct {
	Payload string
}

func (o *SetVcsRootPropertyOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/vcs-roots/{vcsRootLocator}/properties/{name}][%d] setVcsRootPropertyOK  %+v", 200, o.Payload)
}
func (o *SetVcsRootPropertyOK) GetPayload() string {
	return o.Payload
}

func (o *SetVcsRootPropertyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
