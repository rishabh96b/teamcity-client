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

// GetAllSnapshotDependenciesReader is a Reader for the GetAllSnapshotDependencies structure.
type GetAllSnapshotDependenciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllSnapshotDependenciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllSnapshotDependenciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllSnapshotDependenciesOK creates a GetAllSnapshotDependenciesOK with default headers values
func NewGetAllSnapshotDependenciesOK() *GetAllSnapshotDependenciesOK {
	return &GetAllSnapshotDependenciesOK{}
}

/* GetAllSnapshotDependenciesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllSnapshotDependenciesOK struct {
	Payload *models.SnapshotDependencies
}

func (o *GetAllSnapshotDependenciesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/snapshot-dependencies][%d] getAllSnapshotDependenciesOK  %+v", 200, o.Payload)
}
func (o *GetAllSnapshotDependenciesOK) GetPayload() *models.SnapshotDependencies {
	return o.Payload
}

func (o *GetAllSnapshotDependenciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotDependencies)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
