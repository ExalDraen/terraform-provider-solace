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

// GetMsgVpnTopicEndpointsReader is a Reader for the GetMsgVpnTopicEndpoints structure.
type GetMsgVpnTopicEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnTopicEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnTopicEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnTopicEndpointsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnTopicEndpointsOK creates a GetMsgVpnTopicEndpointsOK with default headers values
func NewGetMsgVpnTopicEndpointsOK() *GetMsgVpnTopicEndpointsOK {
	return &GetMsgVpnTopicEndpointsOK{}
}

/*GetMsgVpnTopicEndpointsOK handles this case with default header values.

The list of Topic Endpoint objects' attributes, and the request metadata.
*/
type GetMsgVpnTopicEndpointsOK struct {
	Payload *models.MsgVpnTopicEndpointsResponse
}

func (o *GetMsgVpnTopicEndpointsOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/topicEndpoints][%d] getMsgVpnTopicEndpointsOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnTopicEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnTopicEndpointsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnTopicEndpointsDefault creates a GetMsgVpnTopicEndpointsDefault with default headers values
func NewGetMsgVpnTopicEndpointsDefault(code int) *GetMsgVpnTopicEndpointsDefault {
	return &GetMsgVpnTopicEndpointsDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnTopicEndpointsDefault handles this case with default header values.

Error response
*/
type GetMsgVpnTopicEndpointsDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn topic endpoints default response
func (o *GetMsgVpnTopicEndpointsDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnTopicEndpointsDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/topicEndpoints][%d] getMsgVpnTopicEndpoints default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnTopicEndpointsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}