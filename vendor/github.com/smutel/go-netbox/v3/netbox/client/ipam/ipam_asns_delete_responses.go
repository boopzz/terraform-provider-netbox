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
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamAsnsDeleteReader is a Reader for the IpamAsnsDelete structure.
type IpamAsnsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAsnsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamAsnsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamAsnsDeleteNoContent creates a IpamAsnsDeleteNoContent with default headers values
func NewIpamAsnsDeleteNoContent() *IpamAsnsDeleteNoContent {
	return &IpamAsnsDeleteNoContent{}
}

/*
IpamAsnsDeleteNoContent describes a response with status code 204, with default header values.

IpamAsnsDeleteNoContent ipam asns delete no content
*/
type IpamAsnsDeleteNoContent struct {
}

// IsSuccess returns true when this ipam asns delete no content response has a 2xx status code
func (o *IpamAsnsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam asns delete no content response has a 3xx status code
func (o *IpamAsnsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam asns delete no content response has a 4xx status code
func (o *IpamAsnsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam asns delete no content response has a 5xx status code
func (o *IpamAsnsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam asns delete no content response a status code equal to that given
func (o *IpamAsnsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamAsnsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/asns/{id}/][%d] ipamAsnsDeleteNoContent ", 204)
}

func (o *IpamAsnsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/asns/{id}/][%d] ipamAsnsDeleteNoContent ", 204)
}

func (o *IpamAsnsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
