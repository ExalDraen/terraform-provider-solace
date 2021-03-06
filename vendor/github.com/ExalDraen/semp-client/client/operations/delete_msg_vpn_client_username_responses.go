// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ExalDraen/semp-client/models"
)

// DeleteMsgVpnClientUsernameReader is a Reader for the DeleteMsgVpnClientUsername structure.
type DeleteMsgVpnClientUsernameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnClientUsernameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMsgVpnClientUsernameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMsgVpnClientUsernameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnClientUsernameOK creates a DeleteMsgVpnClientUsernameOK with default headers values
func NewDeleteMsgVpnClientUsernameOK() *DeleteMsgVpnClientUsernameOK {
	return &DeleteMsgVpnClientUsernameOK{}
}

/*DeleteMsgVpnClientUsernameOK handles this case with default header values.

The request metadata.
*/
type DeleteMsgVpnClientUsernameOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnClientUsernameOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/clientUsernames/{clientUsername}][%d] deleteMsgVpnClientUsernameOK  %+v", 200, o.Payload)
}

func (o *DeleteMsgVpnClientUsernameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnClientUsernameDefault creates a DeleteMsgVpnClientUsernameDefault with default headers values
func NewDeleteMsgVpnClientUsernameDefault(code int) *DeleteMsgVpnClientUsernameDefault {
	return &DeleteMsgVpnClientUsernameDefault{
		_statusCode: code,
	}
}

/*DeleteMsgVpnClientUsernameDefault handles this case with default header values.

Error response
*/
type DeleteMsgVpnClientUsernameDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn client username default response
func (o *DeleteMsgVpnClientUsernameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnClientUsernameDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/clientUsernames/{clientUsername}][%d] deleteMsgVpnClientUsername default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMsgVpnClientUsernameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
