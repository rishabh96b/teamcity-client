// Code generated by go-swagger; DO NOT EDIT.

package change

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// GetChangeFirstBuildsReader is a Reader for the GetChangeFirstBuilds structure.
type GetChangeFirstBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChangeFirstBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetChangeFirstBuildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetChangeFirstBuildsOK creates a GetChangeFirstBuildsOK with default headers values
func NewGetChangeFirstBuildsOK() *GetChangeFirstBuildsOK {
	return &GetChangeFirstBuildsOK{}
}

/* GetChangeFirstBuildsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetChangeFirstBuildsOK struct {
	Payload *models.Builds
}

func (o *GetChangeFirstBuildsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/changes/{changeLocator}/firstBuilds][%d] getChangeFirstBuildsOK  %+v", 200, o.Payload)
}
func (o *GetChangeFirstBuildsOK) GetPayload() *models.Builds {
	return o.Payload
}

func (o *GetChangeFirstBuildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Builds)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
