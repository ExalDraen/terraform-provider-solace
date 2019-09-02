// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteMsgVpnMqttSessionSubscriptionParams creates a new DeleteMsgVpnMqttSessionSubscriptionParams object
// with the default values initialized.
func NewDeleteMsgVpnMqttSessionSubscriptionParams() *DeleteMsgVpnMqttSessionSubscriptionParams {
	var ()
	return &DeleteMsgVpnMqttSessionSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnMqttSessionSubscriptionParamsWithTimeout creates a new DeleteMsgVpnMqttSessionSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMsgVpnMqttSessionSubscriptionParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnMqttSessionSubscriptionParams {
	var ()
	return &DeleteMsgVpnMqttSessionSubscriptionParams{

		timeout: timeout,
	}
}

// NewDeleteMsgVpnMqttSessionSubscriptionParamsWithContext creates a new DeleteMsgVpnMqttSessionSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMsgVpnMqttSessionSubscriptionParamsWithContext(ctx context.Context) *DeleteMsgVpnMqttSessionSubscriptionParams {
	var ()
	return &DeleteMsgVpnMqttSessionSubscriptionParams{

		Context: ctx,
	}
}

// NewDeleteMsgVpnMqttSessionSubscriptionParamsWithHTTPClient creates a new DeleteMsgVpnMqttSessionSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMsgVpnMqttSessionSubscriptionParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnMqttSessionSubscriptionParams {
	var ()
	return &DeleteMsgVpnMqttSessionSubscriptionParams{
		HTTPClient: client,
	}
}

/*DeleteMsgVpnMqttSessionSubscriptionParams contains all the parameters to send to the API endpoint
for the delete msg vpn mqtt session subscription operation typically these are written to a http.Request
*/
type DeleteMsgVpnMqttSessionSubscriptionParams struct {

	/*MqttSessionClientID
	  The mqttSessionClientId of the MQTT Session.

	*/
	MqttSessionClientID string
	/*MqttSessionVirtualRouter
	  The mqttSessionVirtualRouter of the MQTT Session.

	*/
	MqttSessionVirtualRouter string
	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*SubscriptionTopic
	  The subscriptionTopic of the MQTT Session Subscription.

	*/
	SubscriptionTopic string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnMqttSessionSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) WithContext(ctx context.Context) *DeleteMsgVpnMqttSessionSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnMqttSessionSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMqttSessionClientID adds the mqttSessionClientID to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) WithMqttSessionClientID(mqttSessionClientID string) *DeleteMsgVpnMqttSessionSubscriptionParams {
	o.SetMqttSessionClientID(mqttSessionClientID)
	return o
}

// SetMqttSessionClientID adds the mqttSessionClientId to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) SetMqttSessionClientID(mqttSessionClientID string) {
	o.MqttSessionClientID = mqttSessionClientID
}

// WithMqttSessionVirtualRouter adds the mqttSessionVirtualRouter to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) WithMqttSessionVirtualRouter(mqttSessionVirtualRouter string) *DeleteMsgVpnMqttSessionSubscriptionParams {
	o.SetMqttSessionVirtualRouter(mqttSessionVirtualRouter)
	return o
}

// SetMqttSessionVirtualRouter adds the mqttSessionVirtualRouter to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) SetMqttSessionVirtualRouter(mqttSessionVirtualRouter string) {
	o.MqttSessionVirtualRouter = mqttSessionVirtualRouter
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnMqttSessionSubscriptionParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSubscriptionTopic adds the subscriptionTopic to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) WithSubscriptionTopic(subscriptionTopic string) *DeleteMsgVpnMqttSessionSubscriptionParams {
	o.SetSubscriptionTopic(subscriptionTopic)
	return o
}

// SetSubscriptionTopic adds the subscriptionTopic to the delete msg vpn mqtt session subscription params
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) SetSubscriptionTopic(subscriptionTopic string) {
	o.SubscriptionTopic = subscriptionTopic
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnMqttSessionSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mqttSessionClientId
	if err := r.SetPathParam("mqttSessionClientId", o.MqttSessionClientID); err != nil {
		return err
	}

	// path param mqttSessionVirtualRouter
	if err := r.SetPathParam("mqttSessionVirtualRouter", o.MqttSessionVirtualRouter); err != nil {
		return err
	}

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	// path param subscriptionTopic
	if err := r.SetPathParam("subscriptionTopic", o.SubscriptionTopic); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}