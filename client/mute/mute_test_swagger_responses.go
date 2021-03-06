// Code generated by go-swagger; DO NOT EDIT.

package mute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// MuteTestReader is a Reader for the MuteTest structure.
type MuteTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MuteTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMuteTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMuteTestOK creates a MuteTestOK with default headers values
func NewMuteTestOK() *MuteTestOK {
	return &MuteTestOK{}
}

/* MuteTestOK describes a response with status code 200, with default header values.

successful operation
*/
type MuteTestOK struct {
	Payload *models.Mute
}

func (o *MuteTestOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/mutes][%d] muteTestOK  %+v", 200, o.Payload)
}
func (o *MuteTestOK) GetPayload() *models.Mute {
	return o.Payload
}

func (o *MuteTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Mute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
