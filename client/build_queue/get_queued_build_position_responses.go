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

// GetQueuedBuildPositionReader is a Reader for the GetQueuedBuildPosition structure.
type GetQueuedBuildPositionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQueuedBuildPositionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQueuedBuildPositionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQueuedBuildPositionOK creates a GetQueuedBuildPositionOK with default headers values
func NewGetQueuedBuildPositionOK() *GetQueuedBuildPositionOK {
	return &GetQueuedBuildPositionOK{}
}

/* GetQueuedBuildPositionOK describes a response with status code 200, with default header values.

successful operation
*/
type GetQueuedBuildPositionOK struct {
	Payload *models.Build
}

func (o *GetQueuedBuildPositionOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildQueue/order/{queuePosition}][%d] getQueuedBuildPositionOK  %+v", 200, o.Payload)
}
func (o *GetQueuedBuildPositionOK) GetPayload() *models.Build {
	return o.Payload
}

func (o *GetQueuedBuildPositionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Build)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
