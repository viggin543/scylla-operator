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

// NewFindConfigCommitlogReuseSegmentsParams creates a new FindConfigCommitlogReuseSegmentsParams object
// with the default values initialized.
func NewFindConfigCommitlogReuseSegmentsParams() *FindConfigCommitlogReuseSegmentsParams {

	return &FindConfigCommitlogReuseSegmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigCommitlogReuseSegmentsParamsWithTimeout creates a new FindConfigCommitlogReuseSegmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigCommitlogReuseSegmentsParamsWithTimeout(timeout time.Duration) *FindConfigCommitlogReuseSegmentsParams {

	return &FindConfigCommitlogReuseSegmentsParams{

		timeout: timeout,
	}
}

// NewFindConfigCommitlogReuseSegmentsParamsWithContext creates a new FindConfigCommitlogReuseSegmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigCommitlogReuseSegmentsParamsWithContext(ctx context.Context) *FindConfigCommitlogReuseSegmentsParams {

	return &FindConfigCommitlogReuseSegmentsParams{

		Context: ctx,
	}
}

// NewFindConfigCommitlogReuseSegmentsParamsWithHTTPClient creates a new FindConfigCommitlogReuseSegmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigCommitlogReuseSegmentsParamsWithHTTPClient(client *http.Client) *FindConfigCommitlogReuseSegmentsParams {

	return &FindConfigCommitlogReuseSegmentsParams{
		HTTPClient: client,
	}
}

/*
FindConfigCommitlogReuseSegmentsParams contains all the parameters to send to the API endpoint
for the find config commitlog reuse segments operation typically these are written to a http.Request
*/
type FindConfigCommitlogReuseSegmentsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config commitlog reuse segments params
func (o *FindConfigCommitlogReuseSegmentsParams) WithTimeout(timeout time.Duration) *FindConfigCommitlogReuseSegmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config commitlog reuse segments params
func (o *FindConfigCommitlogReuseSegmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config commitlog reuse segments params
func (o *FindConfigCommitlogReuseSegmentsParams) WithContext(ctx context.Context) *FindConfigCommitlogReuseSegmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config commitlog reuse segments params
func (o *FindConfigCommitlogReuseSegmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config commitlog reuse segments params
func (o *FindConfigCommitlogReuseSegmentsParams) WithHTTPClient(client *http.Client) *FindConfigCommitlogReuseSegmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config commitlog reuse segments params
func (o *FindConfigCommitlogReuseSegmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigCommitlogReuseSegmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
