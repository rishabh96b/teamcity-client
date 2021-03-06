// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// GetAllBranchesReader is a Reader for the GetAllBranches structure.
type GetAllBranchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllBranchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllBranchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllBranchesOK creates a GetAllBranchesOK with default headers values
func NewGetAllBranchesOK() *GetAllBranchesOK {
	return &GetAllBranchesOK{}
}

/* GetAllBranchesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllBranchesOK struct {
	Payload *models.Branches
}

func (o *GetAllBranchesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/projects/{projectLocator}/branches][%d] getAllBranchesOK  %+v", 200, o.Payload)
}
func (o *GetAllBranchesOK) GetPayload() *models.Branches {
	return o.Payload
}

func (o *GetAllBranchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Branches)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
