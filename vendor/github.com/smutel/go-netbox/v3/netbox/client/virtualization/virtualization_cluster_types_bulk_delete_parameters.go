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

package virtualization

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

// NewVirtualizationClusterTypesBulkDeleteParams creates a new VirtualizationClusterTypesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationClusterTypesBulkDeleteParams() *VirtualizationClusterTypesBulkDeleteParams {
	return &VirtualizationClusterTypesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterTypesBulkDeleteParamsWithTimeout creates a new VirtualizationClusterTypesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewVirtualizationClusterTypesBulkDeleteParamsWithTimeout(timeout time.Duration) *VirtualizationClusterTypesBulkDeleteParams {
	return &VirtualizationClusterTypesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewVirtualizationClusterTypesBulkDeleteParamsWithContext creates a new VirtualizationClusterTypesBulkDeleteParams object
// with the ability to set a context for a request.
func NewVirtualizationClusterTypesBulkDeleteParamsWithContext(ctx context.Context) *VirtualizationClusterTypesBulkDeleteParams {
	return &VirtualizationClusterTypesBulkDeleteParams{
		Context: ctx,
	}
}

// NewVirtualizationClusterTypesBulkDeleteParamsWithHTTPClient creates a new VirtualizationClusterTypesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationClusterTypesBulkDeleteParamsWithHTTPClient(client *http.Client) *VirtualizationClusterTypesBulkDeleteParams {
	return &VirtualizationClusterTypesBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
VirtualizationClusterTypesBulkDeleteParams contains all the parameters to send to the API endpoint

	for the virtualization cluster types bulk delete operation.

	Typically these are written to a http.Request.
*/
type VirtualizationClusterTypesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization cluster types bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClusterTypesBulkDeleteParams) WithDefaults() *VirtualizationClusterTypesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization cluster types bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClusterTypesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization cluster types bulk delete params
func (o *VirtualizationClusterTypesBulkDeleteParams) WithTimeout(timeout time.Duration) *VirtualizationClusterTypesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster types bulk delete params
func (o *VirtualizationClusterTypesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster types bulk delete params
func (o *VirtualizationClusterTypesBulkDeleteParams) WithContext(ctx context.Context) *VirtualizationClusterTypesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster types bulk delete params
func (o *VirtualizationClusterTypesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster types bulk delete params
func (o *VirtualizationClusterTypesBulkDeleteParams) WithHTTPClient(client *http.Client) *VirtualizationClusterTypesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster types bulk delete params
func (o *VirtualizationClusterTypesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterTypesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
