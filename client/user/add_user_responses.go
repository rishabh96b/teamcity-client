// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// AddUserReader is a Reader for the AddUser structure.
type AddUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddUserOK creates a AddUserOK with default headers values
func NewAddUserOK() *AddUserOK {
	return &AddUserOK{}
}

/* AddUserOK describes a response with status code 200, with default header values.

successful operation
*/
type AddUserOK struct {
	Payload *models.User
}

func (o *AddUserOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/users][%d] addUserOK  %+v", 200, o.Payload)
}
func (o *AddUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *AddUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
