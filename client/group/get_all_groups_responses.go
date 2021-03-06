// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// GetAllGroupsReader is a Reader for the GetAllGroups structure.
type GetAllGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllGroupsOK creates a GetAllGroupsOK with default headers values
func NewGetAllGroupsOK() *GetAllGroupsOK {
	return &GetAllGroupsOK{}
}

/* GetAllGroupsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllGroupsOK struct {
	Payload *models.Groups
}

func (o *GetAllGroupsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/userGroups][%d] getAllGroupsOK  %+v", 200, o.Payload)
}
func (o *GetAllGroupsOK) GetPayload() *models.Groups {
	return o.Payload
}

func (o *GetAllGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Groups)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
