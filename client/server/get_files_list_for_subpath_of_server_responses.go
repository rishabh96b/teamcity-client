// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rishabh96b/teamcity-client/models"
)

// GetFilesListForSubpathOfServerReader is a Reader for the GetFilesListForSubpathOfServer structure.
type GetFilesListForSubpathOfServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilesListForSubpathOfServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFilesListForSubpathOfServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFilesListForSubpathOfServerOK creates a GetFilesListForSubpathOfServerOK with default headers values
func NewGetFilesListForSubpathOfServerOK() *GetFilesListForSubpathOfServerOK {
	return &GetFilesListForSubpathOfServerOK{}
}

/* GetFilesListForSubpathOfServerOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFilesListForSubpathOfServerOK struct {
	Payload *models.Files
}

func (o *GetFilesListForSubpathOfServerOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/server/files/{areaId}/{path}][%d] getFilesListForSubpathOfServerOK  %+v", 200, o.Payload)
}
func (o *GetFilesListForSubpathOfServerOK) GetPayload() *models.Files {
	return o.Payload
}

func (o *GetFilesListForSubpathOfServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Files)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
