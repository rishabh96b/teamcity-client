// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StartBackupReader is a Reader for the StartBackup structure.
type StartBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStartBackupOK creates a StartBackupOK with default headers values
func NewStartBackupOK() *StartBackupOK {
	return &StartBackupOK{}
}

/* StartBackupOK describes a response with status code 200, with default header values.

successful operation
*/
type StartBackupOK struct {
	Payload string
}

func (o *StartBackupOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/server/backup][%d] startBackupOK  %+v", 200, o.Payload)
}
func (o *StartBackupOK) GetPayload() string {
	return o.Payload
}

func (o *StartBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
