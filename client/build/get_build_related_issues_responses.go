// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// GetBuildRelatedIssuesReader is a Reader for the GetBuildRelatedIssues structure.
type GetBuildRelatedIssuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildRelatedIssuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildRelatedIssuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBuildRelatedIssuesOK creates a GetBuildRelatedIssuesOK with default headers values
func NewGetBuildRelatedIssuesOK() *GetBuildRelatedIssuesOK {
	return &GetBuildRelatedIssuesOK{}
}

/* GetBuildRelatedIssuesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBuildRelatedIssuesOK struct {
	Payload *models.IssuesUsages
}

func (o *GetBuildRelatedIssuesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/builds/{buildLocator}/relatedIssues][%d] getBuildRelatedIssuesOK  %+v", 200, o.Payload)
}
func (o *GetBuildRelatedIssuesOK) GetPayload() *models.IssuesUsages {
	return o.Payload
}

func (o *GetBuildRelatedIssuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IssuesUsages)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
