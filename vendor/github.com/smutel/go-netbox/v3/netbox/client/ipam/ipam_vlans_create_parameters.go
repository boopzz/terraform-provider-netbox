// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// NewIpamVlansCreateParams creates a new IpamVlansCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlansCreateParams() *IpamVlansCreateParams {
	return &IpamVlansCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlansCreateParamsWithTimeout creates a new IpamVlansCreateParams object
// with the ability to set a timeout on a request.
func NewIpamVlansCreateParamsWithTimeout(timeout time.Duration) *IpamVlansCreateParams {
	return &IpamVlansCreateParams{
		timeout: timeout,
	}
}

// NewIpamVlansCreateParamsWithContext creates a new IpamVlansCreateParams object
// with the ability to set a context for a request.
func NewIpamVlansCreateParamsWithContext(ctx context.Context) *IpamVlansCreateParams {
	return &IpamVlansCreateParams{
		Context: ctx,
	}
}

// NewIpamVlansCreateParamsWithHTTPClient creates a new IpamVlansCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlansCreateParamsWithHTTPClient(client *http.Client) *IpamVlansCreateParams {
	return &IpamVlansCreateParams{
		HTTPClient: client,
	}
}

/*
IpamVlansCreateParams contains all the parameters to send to the API endpoint

	for the ipam vlans create operation.

	Typically these are written to a http.Request.
*/
type IpamVlansCreateParams struct {

	// Data.
	Data *models.WritableVLAN

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlans create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansCreateParams) WithDefaults() *IpamVlansCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlans create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlans create params
func (o *IpamVlansCreateParams) WithTimeout(timeout time.Duration) *IpamVlansCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans create params
func (o *IpamVlansCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans create params
func (o *IpamVlansCreateParams) WithContext(ctx context.Context) *IpamVlansCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans create params
func (o *IpamVlansCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans create params
func (o *IpamVlansCreateParams) WithHTTPClient(client *http.Client) *IpamVlansCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans create params
func (o *IpamVlansCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlans create params
func (o *IpamVlansCreateParams) WithData(data *models.WritableVLAN) *IpamVlansCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlans create params
func (o *IpamVlansCreateParams) SetData(data *models.WritableVLAN) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlansCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
