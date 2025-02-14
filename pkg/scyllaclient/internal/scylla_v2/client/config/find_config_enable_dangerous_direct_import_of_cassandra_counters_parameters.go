// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewFindConfigEnableDangerousDirectImportOfCassandraCountersParams creates a new FindConfigEnableDangerousDirectImportOfCassandraCountersParams object
// with the default values initialized.
func NewFindConfigEnableDangerousDirectImportOfCassandraCountersParams() *FindConfigEnableDangerousDirectImportOfCassandraCountersParams {

	return &FindConfigEnableDangerousDirectImportOfCassandraCountersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigEnableDangerousDirectImportOfCassandraCountersParamsWithTimeout creates a new FindConfigEnableDangerousDirectImportOfCassandraCountersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigEnableDangerousDirectImportOfCassandraCountersParamsWithTimeout(timeout time.Duration) *FindConfigEnableDangerousDirectImportOfCassandraCountersParams {

	return &FindConfigEnableDangerousDirectImportOfCassandraCountersParams{

		timeout: timeout,
	}
}

// NewFindConfigEnableDangerousDirectImportOfCassandraCountersParamsWithContext creates a new FindConfigEnableDangerousDirectImportOfCassandraCountersParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigEnableDangerousDirectImportOfCassandraCountersParamsWithContext(ctx context.Context) *FindConfigEnableDangerousDirectImportOfCassandraCountersParams {

	return &FindConfigEnableDangerousDirectImportOfCassandraCountersParams{

		Context: ctx,
	}
}

// NewFindConfigEnableDangerousDirectImportOfCassandraCountersParamsWithHTTPClient creates a new FindConfigEnableDangerousDirectImportOfCassandraCountersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigEnableDangerousDirectImportOfCassandraCountersParamsWithHTTPClient(client *http.Client) *FindConfigEnableDangerousDirectImportOfCassandraCountersParams {

	return &FindConfigEnableDangerousDirectImportOfCassandraCountersParams{
		HTTPClient: client,
	}
}

/*
FindConfigEnableDangerousDirectImportOfCassandraCountersParams contains all the parameters to send to the API endpoint
for the find config enable dangerous direct import of cassandra counters operation typically these are written to a http.Request
*/
type FindConfigEnableDangerousDirectImportOfCassandraCountersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config enable dangerous direct import of cassandra counters params
func (o *FindConfigEnableDangerousDirectImportOfCassandraCountersParams) WithTimeout(timeout time.Duration) *FindConfigEnableDangerousDirectImportOfCassandraCountersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config enable dangerous direct import of cassandra counters params
func (o *FindConfigEnableDangerousDirectImportOfCassandraCountersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config enable dangerous direct import of cassandra counters params
func (o *FindConfigEnableDangerousDirectImportOfCassandraCountersParams) WithContext(ctx context.Context) *FindConfigEnableDangerousDirectImportOfCassandraCountersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config enable dangerous direct import of cassandra counters params
func (o *FindConfigEnableDangerousDirectImportOfCassandraCountersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config enable dangerous direct import of cassandra counters params
func (o *FindConfigEnableDangerousDirectImportOfCassandraCountersParams) WithHTTPClient(client *http.Client) *FindConfigEnableDangerousDirectImportOfCassandraCountersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config enable dangerous direct import of cassandra counters params
func (o *FindConfigEnableDangerousDirectImportOfCassandraCountersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigEnableDangerousDirectImportOfCassandraCountersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
