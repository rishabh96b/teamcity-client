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

// GetAllUserGroupsReader is a Reader for the GetAllUserGroups structure.
type GetAllUserGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllUserGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllUserGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllUserGroupsOK creates a GetAllUserGroupsOK with default headers values
func NewGetAllUserGroupsOK() *GetAllUserGroupsOK {
	return &GetAllUserGroupsOK{}
}

/* GetAllUserGroupsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllUserGroupsOK struct {
	Payload *models.Groups
}

func (o *GetAllUserGroupsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/users/{userLocator}/groups][%d] getAllUserGroupsOK  %+v", 200, o.Payload)
}
func (o *GetAllUserGroupsOK) GetPayload() *models.Groups {
	return o.Payload
}

func (o *GetAllUserGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Groups)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
