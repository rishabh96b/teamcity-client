// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// GetAllVcsRootsOfBuildTypeReader is a Reader for the GetAllVcsRootsOfBuildType structure.
type GetAllVcsRootsOfBuildTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllVcsRootsOfBuildTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllVcsRootsOfBuildTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllVcsRootsOfBuildTypeOK creates a GetAllVcsRootsOfBuildTypeOK with default headers values
func NewGetAllVcsRootsOfBuildTypeOK() *GetAllVcsRootsOfBuildTypeOK {
	return &GetAllVcsRootsOfBuildTypeOK{}
}

/* GetAllVcsRootsOfBuildTypeOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllVcsRootsOfBuildTypeOK struct {
	Payload *models.VcsRootEntries
}

func (o *GetAllVcsRootsOfBuildTypeOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/vcs-root-entries][%d] getAllVcsRootsOfBuildTypeOK  %+v", 200, o.Payload)
}
func (o *GetAllVcsRootsOfBuildTypeOK) GetPayload() *models.VcsRootEntries {
	return o.Payload
}

func (o *GetAllVcsRootsOfBuildTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsRootEntries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
