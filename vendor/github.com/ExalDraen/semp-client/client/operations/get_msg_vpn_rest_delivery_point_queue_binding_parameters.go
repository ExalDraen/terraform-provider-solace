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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMsgVpnRestDeliveryPointQueueBindingParams creates a new GetMsgVpnRestDeliveryPointQueueBindingParams object
// with the default values initialized.
func NewGetMsgVpnRestDeliveryPointQueueBindingParams() *GetMsgVpnRestDeliveryPointQueueBindingParams {
	var ()
	return &GetMsgVpnRestDeliveryPointQueueBindingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnRestDeliveryPointQueueBindingParamsWithTimeout creates a new GetMsgVpnRestDeliveryPointQueueBindingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnRestDeliveryPointQueueBindingParamsWithTimeout(timeout time.Duration) *GetMsgVpnRestDeliveryPointQueueBindingParams {
	var ()
	return &GetMsgVpnRestDeliveryPointQueueBindingParams{

		timeout: timeout,
	}
}

// NewGetMsgVpnRestDeliveryPointQueueBindingParamsWithContext creates a new GetMsgVpnRestDeliveryPointQueueBindingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnRestDeliveryPointQueueBindingParamsWithContext(ctx context.Context) *GetMsgVpnRestDeliveryPointQueueBindingParams {
	var ()
	return &GetMsgVpnRestDeliveryPointQueueBindingParams{

		Context: ctx,
	}
}

// NewGetMsgVpnRestDeliveryPointQueueBindingParamsWithHTTPClient creates a new GetMsgVpnRestDeliveryPointQueueBindingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnRestDeliveryPointQueueBindingParamsWithHTTPClient(client *http.Client) *GetMsgVpnRestDeliveryPointQueueBindingParams {
	var ()
	return &GetMsgVpnRestDeliveryPointQueueBindingParams{
		HTTPClient: client,
	}
}

/*GetMsgVpnRestDeliveryPointQueueBindingParams contains all the parameters to send to the API endpoint
for the get msg vpn rest delivery point queue binding operation typically these are written to a http.Request
*/
type GetMsgVpnRestDeliveryPointQueueBindingParams struct {

	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*QueueBindingName
	  The queueBindingName of the Queue Binding.

	*/
	QueueBindingName string
	/*RestDeliveryPointName
	  The restDeliveryPointName of the REST Delivery Point.

	*/
	RestDeliveryPointName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) WithTimeout(timeout time.Duration) *GetMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) WithContext(ctx context.Context) *GetMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) WithHTTPClient(client *http.Client) *GetMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithQueueBindingName adds the queueBindingName to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) WithQueueBindingName(queueBindingName string) *GetMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetQueueBindingName(queueBindingName)
	return o
}

// SetQueueBindingName adds the queueBindingName to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) SetQueueBindingName(queueBindingName string) {
	o.QueueBindingName = queueBindingName
}

// WithRestDeliveryPointName adds the restDeliveryPointName to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) WithRestDeliveryPointName(restDeliveryPointName string) *GetMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetRestDeliveryPointName(restDeliveryPointName)
	return o
}

// SetRestDeliveryPointName adds the restDeliveryPointName to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) SetRestDeliveryPointName(restDeliveryPointName string) {
	o.RestDeliveryPointName = restDeliveryPointName
}

// WithSelect adds the selectVar to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) WithSelect(selectVar []string) *GetMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn rest delivery point queue binding params
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnRestDeliveryPointQueueBindingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	// path param queueBindingName
	if err := r.SetPathParam("queueBindingName", o.QueueBindingName); err != nil {
		return err
	}

	// path param restDeliveryPointName
	if err := r.SetPathParam("restDeliveryPointName", o.RestDeliveryPointName); err != nil {
		return err
	}

	valuesSelect := o.Select

	joinedSelect := swag.JoinByFormat(valuesSelect, "csv")
	// query array param select
	if err := r.SetQueryParam("select", joinedSelect...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
