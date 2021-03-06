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

// NewDeleteMsgVpnJndiTopicParams creates a new DeleteMsgVpnJndiTopicParams object
// with the default values initialized.
func NewDeleteMsgVpnJndiTopicParams() *DeleteMsgVpnJndiTopicParams {
	var ()
	return &DeleteMsgVpnJndiTopicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnJndiTopicParamsWithTimeout creates a new DeleteMsgVpnJndiTopicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMsgVpnJndiTopicParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnJndiTopicParams {
	var ()
	return &DeleteMsgVpnJndiTopicParams{

		timeout: timeout,
	}
}

// NewDeleteMsgVpnJndiTopicParamsWithContext creates a new DeleteMsgVpnJndiTopicParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMsgVpnJndiTopicParamsWithContext(ctx context.Context) *DeleteMsgVpnJndiTopicParams {
	var ()
	return &DeleteMsgVpnJndiTopicParams{

		Context: ctx,
	}
}

// NewDeleteMsgVpnJndiTopicParamsWithHTTPClient creates a new DeleteMsgVpnJndiTopicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMsgVpnJndiTopicParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnJndiTopicParams {
	var ()
	return &DeleteMsgVpnJndiTopicParams{
		HTTPClient: client,
	}
}

/*DeleteMsgVpnJndiTopicParams contains all the parameters to send to the API endpoint
for the delete msg vpn jndi topic operation typically these are written to a http.Request
*/
type DeleteMsgVpnJndiTopicParams struct {

	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*TopicName
	  The topicName of the JNDI Topic.

	*/
	TopicName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete msg vpn jndi topic params
func (o *DeleteMsgVpnJndiTopicParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnJndiTopicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn jndi topic params
func (o *DeleteMsgVpnJndiTopicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn jndi topic params
func (o *DeleteMsgVpnJndiTopicParams) WithContext(ctx context.Context) *DeleteMsgVpnJndiTopicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn jndi topic params
func (o *DeleteMsgVpnJndiTopicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn jndi topic params
func (o *DeleteMsgVpnJndiTopicParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnJndiTopicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn jndi topic params
func (o *DeleteMsgVpnJndiTopicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn jndi topic params
func (o *DeleteMsgVpnJndiTopicParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnJndiTopicParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn jndi topic params
func (o *DeleteMsgVpnJndiTopicParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithTopicName adds the topicName to the delete msg vpn jndi topic params
func (o *DeleteMsgVpnJndiTopicParams) WithTopicName(topicName string) *DeleteMsgVpnJndiTopicParams {
	o.SetTopicName(topicName)
	return o
}

// SetTopicName adds the topicName to the delete msg vpn jndi topic params
func (o *DeleteMsgVpnJndiTopicParams) SetTopicName(topicName string) {
	o.TopicName = topicName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnJndiTopicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	// path param topicName
	if err := r.SetPathParam("topicName", o.TopicName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
