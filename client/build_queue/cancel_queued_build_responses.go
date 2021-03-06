// Code generated by go-swagger; DO NOT EDIT.

package build_queue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// CancelQueuedBuildReader is a Reader for the CancelQueuedBuild structure.
type CancelQueuedBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelQueuedBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelQueuedBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelQueuedBuildOK creates a CancelQueuedBuildOK with default headers values
func NewCancelQueuedBuildOK() *CancelQueuedBuildOK {
	return &CancelQueuedBuildOK{}
}

/* CancelQueuedBuildOK describes a response with status code 200, with default header values.

successful operation
*/
type CancelQueuedBuildOK struct {
	Payload *models.Build
}

func (o *CancelQueuedBuildOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/buildQueue/{queuedBuildLocator}][%d] cancelQueuedBuildOK  %+v", 200, o.Payload)
}
func (o *CancelQueuedBuildOK) GetPayload() *models.Build {
	return o.Payload
}

func (o *CancelQueuedBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Build)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
