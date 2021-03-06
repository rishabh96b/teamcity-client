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

// SetMultipleBuildCommentsReader is a Reader for the SetMultipleBuildComments structure.
type SetMultipleBuildCommentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetMultipleBuildCommentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetMultipleBuildCommentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetMultipleBuildCommentsOK creates a SetMultipleBuildCommentsOK with default headers values
func NewSetMultipleBuildCommentsOK() *SetMultipleBuildCommentsOK {
	return &SetMultipleBuildCommentsOK{}
}

/* SetMultipleBuildCommentsOK describes a response with status code 200, with default header values.

successful operation
*/
type SetMultipleBuildCommentsOK struct {
	Payload *models.MultipleOperationResult
}

func (o *SetMultipleBuildCommentsOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/builds/multiple/{buildLocator}/comment][%d] setMultipleBuildCommentsOK  %+v", 200, o.Payload)
}
func (o *SetMultipleBuildCommentsOK) GetPayload() *models.MultipleOperationResult {
	return o.Payload
}

func (o *SetMultipleBuildCommentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MultipleOperationResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
