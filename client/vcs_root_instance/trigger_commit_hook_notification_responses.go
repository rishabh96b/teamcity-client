// Code generated by go-swagger; DO NOT EDIT.

package vcs_root_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TriggerCommitHookNotificationReader is a Reader for the TriggerCommitHookNotification structure.
type TriggerCommitHookNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggerCommitHookNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewTriggerCommitHookNotificationDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewTriggerCommitHookNotificationDefault creates a TriggerCommitHookNotificationDefault with default headers values
func NewTriggerCommitHookNotificationDefault(code int) *TriggerCommitHookNotificationDefault {
	return &TriggerCommitHookNotificationDefault{
		_statusCode: code,
	}
}

/* TriggerCommitHookNotificationDefault describes a response with status code -1, with default header values.

successful operation
*/
type TriggerCommitHookNotificationDefault struct {
	_statusCode int
}

// Code gets the status code for the trigger commit hook notification default response
func (o *TriggerCommitHookNotificationDefault) Code() int {
	return o._statusCode
}

func (o *TriggerCommitHookNotificationDefault) Error() string {
	return fmt.Sprintf("[POST /app/rest/vcs-root-instances/commitHookNotification][%d] triggerCommitHookNotification default ", o._statusCode)
}

func (o *TriggerCommitHookNotificationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
