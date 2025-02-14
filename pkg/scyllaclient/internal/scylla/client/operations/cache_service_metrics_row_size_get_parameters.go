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

// NewCacheServiceMetricsRowSizeGetParams creates a new CacheServiceMetricsRowSizeGetParams object
// with the default values initialized.
func NewCacheServiceMetricsRowSizeGetParams() *CacheServiceMetricsRowSizeGetParams {

	return &CacheServiceMetricsRowSizeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsRowSizeGetParamsWithTimeout creates a new CacheServiceMetricsRowSizeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceMetricsRowSizeGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsRowSizeGetParams {

	return &CacheServiceMetricsRowSizeGetParams{

		timeout: timeout,
	}
}

// NewCacheServiceMetricsRowSizeGetParamsWithContext creates a new CacheServiceMetricsRowSizeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceMetricsRowSizeGetParamsWithContext(ctx context.Context) *CacheServiceMetricsRowSizeGetParams {

	return &CacheServiceMetricsRowSizeGetParams{

		Context: ctx,
	}
}

// NewCacheServiceMetricsRowSizeGetParamsWithHTTPClient creates a new CacheServiceMetricsRowSizeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceMetricsRowSizeGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsRowSizeGetParams {

	return &CacheServiceMetricsRowSizeGetParams{
		HTTPClient: client,
	}
}

/*
CacheServiceMetricsRowSizeGetParams contains all the parameters to send to the API endpoint
for the cache service metrics row size get operation typically these are written to a http.Request
*/
type CacheServiceMetricsRowSizeGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service metrics row size get params
func (o *CacheServiceMetricsRowSizeGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsRowSizeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics row size get params
func (o *CacheServiceMetricsRowSizeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics row size get params
func (o *CacheServiceMetricsRowSizeGetParams) WithContext(ctx context.Context) *CacheServiceMetricsRowSizeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics row size get params
func (o *CacheServiceMetricsRowSizeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics row size get params
func (o *CacheServiceMetricsRowSizeGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsRowSizeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics row size get params
func (o *CacheServiceMetricsRowSizeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsRowSizeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
