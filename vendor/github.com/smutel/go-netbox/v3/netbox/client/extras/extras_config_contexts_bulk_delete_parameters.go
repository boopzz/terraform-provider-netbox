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

package extras

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
)

// NewExtrasConfigContextsBulkDeleteParams creates a new ExtrasConfigContextsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextsBulkDeleteParams() *ExtrasConfigContextsBulkDeleteParams {
	return &ExtrasConfigContextsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextsBulkDeleteParamsWithTimeout creates a new ExtrasConfigContextsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextsBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextsBulkDeleteParams {
	return &ExtrasConfigContextsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextsBulkDeleteParamsWithContext creates a new ExtrasConfigContextsBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextsBulkDeleteParamsWithContext(ctx context.Context) *ExtrasConfigContextsBulkDeleteParams {
	return &ExtrasConfigContextsBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextsBulkDeleteParamsWithHTTPClient creates a new ExtrasConfigContextsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextsBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextsBulkDeleteParams {
	return &ExtrasConfigContextsBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
ExtrasConfigContextsBulkDeleteParams contains all the parameters to send to the API endpoint

	for the extras config contexts bulk delete operation.

	Typically these are written to a http.Request.
*/
type ExtrasConfigContextsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config contexts bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsBulkDeleteParams) WithDefaults() *ExtrasConfigContextsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config contexts bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config contexts bulk delete params
func (o *ExtrasConfigContextsBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config contexts bulk delete params
func (o *ExtrasConfigContextsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config contexts bulk delete params
func (o *ExtrasConfigContextsBulkDeleteParams) WithContext(ctx context.Context) *ExtrasConfigContextsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config contexts bulk delete params
func (o *ExtrasConfigContextsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config contexts bulk delete params
func (o *ExtrasConfigContextsBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config contexts bulk delete params
func (o *ExtrasConfigContextsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
