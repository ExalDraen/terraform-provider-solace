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

// NewDeleteMsgVpnRestDeliveryPointParams creates a new DeleteMsgVpnRestDeliveryPointParams object
// with the default values initialized.
func NewDeleteMsgVpnRestDeliveryPointParams() *DeleteMsgVpnRestDeliveryPointParams {
	var ()
	return &DeleteMsgVpnRestDeliveryPointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnRestDeliveryPointParamsWithTimeout creates a new DeleteMsgVpnRestDeliveryPointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMsgVpnRestDeliveryPointParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnRestDeliveryPointParams {
	var ()
	return &DeleteMsgVpnRestDeliveryPointParams{

		timeout: timeout,
	}
}

// NewDeleteMsgVpnRestDeliveryPointParamsWithContext creates a new DeleteMsgVpnRestDeliveryPointParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMsgVpnRestDeliveryPointParamsWithContext(ctx context.Context) *DeleteMsgVpnRestDeliveryPointParams {
	var ()
	return &DeleteMsgVpnRestDeliveryPointParams{

		Context: ctx,
	}
}

// NewDeleteMsgVpnRestDeliveryPointParamsWithHTTPClient creates a new DeleteMsgVpnRestDeliveryPointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMsgVpnRestDeliveryPointParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnRestDeliveryPointParams {
	var ()
	return &DeleteMsgVpnRestDeliveryPointParams{
		HTTPClient: client,
	}
}

/*DeleteMsgVpnRestDeliveryPointParams contains all the parameters to send to the API endpoint
for the delete msg vpn rest delivery point operation typically these are written to a http.Request
*/
type DeleteMsgVpnRestDeliveryPointParams struct {

	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*RestDeliveryPointName
	  The restDeliveryPointName of the REST Delivery Point.

	*/
	RestDeliveryPointName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete msg vpn rest delivery point params
func (o *DeleteMsgVpnRestDeliveryPointParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnRestDeliveryPointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn rest delivery point params
func (o *DeleteMsgVpnRestDeliveryPointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn rest delivery point params
func (o *DeleteMsgVpnRestDeliveryPointParams) WithContext(ctx context.Context) *DeleteMsgVpnRestDeliveryPointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn rest delivery point params
func (o *DeleteMsgVpnRestDeliveryPointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn rest delivery point params
func (o *DeleteMsgVpnRestDeliveryPointParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnRestDeliveryPointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn rest delivery point params
func (o *DeleteMsgVpnRestDeliveryPointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn rest delivery point params
func (o *DeleteMsgVpnRestDeliveryPointParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnRestDeliveryPointParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn rest delivery point params
func (o *DeleteMsgVpnRestDeliveryPointParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithRestDeliveryPointName adds the restDeliveryPointName to the delete msg vpn rest delivery point params
func (o *DeleteMsgVpnRestDeliveryPointParams) WithRestDeliveryPointName(restDeliveryPointName string) *DeleteMsgVpnRestDeliveryPointParams {
	o.SetRestDeliveryPointName(restDeliveryPointName)
	return o
}

// SetRestDeliveryPointName adds the restDeliveryPointName to the delete msg vpn rest delivery point params
func (o *DeleteMsgVpnRestDeliveryPointParams) SetRestDeliveryPointName(restDeliveryPointName string) {
	o.RestDeliveryPointName = restDeliveryPointName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnRestDeliveryPointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	// path param restDeliveryPointName
	if err := r.SetPathParam("restDeliveryPointName", o.RestDeliveryPointName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
