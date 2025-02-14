// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewColumnFamilyMaximumCompactionByNameGetParams creates a new ColumnFamilyMaximumCompactionByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMaximumCompactionByNameGetParams() *ColumnFamilyMaximumCompactionByNameGetParams {
	var ()
	return &ColumnFamilyMaximumCompactionByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMaximumCompactionByNameGetParamsWithTimeout creates a new ColumnFamilyMaximumCompactionByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMaximumCompactionByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMaximumCompactionByNameGetParams {
	var ()
	return &ColumnFamilyMaximumCompactionByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMaximumCompactionByNameGetParamsWithContext creates a new ColumnFamilyMaximumCompactionByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMaximumCompactionByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMaximumCompactionByNameGetParams {
	var ()
	return &ColumnFamilyMaximumCompactionByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMaximumCompactionByNameGetParamsWithHTTPClient creates a new ColumnFamilyMaximumCompactionByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMaximumCompactionByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMaximumCompactionByNameGetParams {
	var ()
	return &ColumnFamilyMaximumCompactionByNameGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMaximumCompactionByNameGetParams contains all the parameters to send to the API endpoint
for the column family maximum compaction by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMaximumCompactionByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family maximum compaction by name get params
func (o *ColumnFamilyMaximumCompactionByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMaximumCompactionByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family maximum compaction by name get params
func (o *ColumnFamilyMaximumCompactionByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family maximum compaction by name get params
func (o *ColumnFamilyMaximumCompactionByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMaximumCompactionByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family maximum compaction by name get params
func (o *ColumnFamilyMaximumCompactionByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family maximum compaction by name get params
func (o *ColumnFamilyMaximumCompactionByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMaximumCompactionByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family maximum compaction by name get params
func (o *ColumnFamilyMaximumCompactionByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family maximum compaction by name get params
func (o *ColumnFamilyMaximumCompactionByNameGetParams) WithName(name string) *ColumnFamilyMaximumCompactionByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family maximum compaction by name get params
func (o *ColumnFamilyMaximumCompactionByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMaximumCompactionByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
