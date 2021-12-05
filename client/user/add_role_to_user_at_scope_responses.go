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

// AddRoleToUserAtScopeReader is a Reader for the AddRoleToUserAtScope structure.
type AddRoleToUserAtScopeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRoleToUserAtScopeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddRoleToUserAtScopeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddRoleToUserAtScopeOK creates a AddRoleToUserAtScopeOK with default headers values
func NewAddRoleToUserAtScopeOK() *AddRoleToUserAtScopeOK {
	return &AddRoleToUserAtScopeOK{}
}

/* AddRoleToUserAtScopeOK describes a response with status code 200, with default header values.

successful operation
*/
type AddRoleToUserAtScopeOK struct {
	Payload *models.Role
}

func (o *AddRoleToUserAtScopeOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/users/{userLocator}/roles/{roleId}/{scope}][%d] addRoleToUserAtScopeOK  %+v", 200, o.Payload)
}
func (o *AddRoleToUserAtScopeOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *AddRoleToUserAtScopeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
