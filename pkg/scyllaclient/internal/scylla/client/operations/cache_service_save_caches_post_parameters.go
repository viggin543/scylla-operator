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

// NewCacheServiceSaveCachesPostParams creates a new CacheServiceSaveCachesPostParams object
// with the default values initialized.
func NewCacheServiceSaveCachesPostParams() *CacheServiceSaveCachesPostParams {

	return &CacheServiceSaveCachesPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceSaveCachesPostParamsWithTimeout creates a new CacheServiceSaveCachesPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceSaveCachesPostParamsWithTimeout(timeout time.Duration) *CacheServiceSaveCachesPostParams {

	return &CacheServiceSaveCachesPostParams{

		timeout: timeout,
	}
}

// NewCacheServiceSaveCachesPostParamsWithContext creates a new CacheServiceSaveCachesPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceSaveCachesPostParamsWithContext(ctx context.Context) *CacheServiceSaveCachesPostParams {

	return &CacheServiceSaveCachesPostParams{

		Context: ctx,
	}
}

// NewCacheServiceSaveCachesPostParamsWithHTTPClient creates a new CacheServiceSaveCachesPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceSaveCachesPostParamsWithHTTPClient(client *http.Client) *CacheServiceSaveCachesPostParams {

	return &CacheServiceSaveCachesPostParams{
		HTTPClient: client,
	}
}

/*
CacheServiceSaveCachesPostParams contains all the parameters to send to the API endpoint
for the cache service save caches post operation typically these are written to a http.Request
*/
type CacheServiceSaveCachesPostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service save caches post params
func (o *CacheServiceSaveCachesPostParams) WithTimeout(timeout time.Duration) *CacheServiceSaveCachesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service save caches post params
func (o *CacheServiceSaveCachesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service save caches post params
func (o *CacheServiceSaveCachesPostParams) WithContext(ctx context.Context) *CacheServiceSaveCachesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service save caches post params
func (o *CacheServiceSaveCachesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service save caches post params
func (o *CacheServiceSaveCachesPostParams) WithHTTPClient(client *http.Client) *CacheServiceSaveCachesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service save caches post params
func (o *CacheServiceSaveCachesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceSaveCachesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
