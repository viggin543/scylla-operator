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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCacheServiceRowCacheSavePeriodPostParams creates a new CacheServiceRowCacheSavePeriodPostParams object
// with the default values initialized.
func NewCacheServiceRowCacheSavePeriodPostParams() *CacheServiceRowCacheSavePeriodPostParams {
	var ()
	return &CacheServiceRowCacheSavePeriodPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceRowCacheSavePeriodPostParamsWithTimeout creates a new CacheServiceRowCacheSavePeriodPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceRowCacheSavePeriodPostParamsWithTimeout(timeout time.Duration) *CacheServiceRowCacheSavePeriodPostParams {
	var ()
	return &CacheServiceRowCacheSavePeriodPostParams{

		timeout: timeout,
	}
}

// NewCacheServiceRowCacheSavePeriodPostParamsWithContext creates a new CacheServiceRowCacheSavePeriodPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceRowCacheSavePeriodPostParamsWithContext(ctx context.Context) *CacheServiceRowCacheSavePeriodPostParams {
	var ()
	return &CacheServiceRowCacheSavePeriodPostParams{

		Context: ctx,
	}
}

// NewCacheServiceRowCacheSavePeriodPostParamsWithHTTPClient creates a new CacheServiceRowCacheSavePeriodPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceRowCacheSavePeriodPostParamsWithHTTPClient(client *http.Client) *CacheServiceRowCacheSavePeriodPostParams {
	var ()
	return &CacheServiceRowCacheSavePeriodPostParams{
		HTTPClient: client,
	}
}

/*
CacheServiceRowCacheSavePeriodPostParams contains all the parameters to send to the API endpoint
for the cache service row cache save period post operation typically these are written to a http.Request
*/
type CacheServiceRowCacheSavePeriodPostParams struct {

	/*Period
	  row cache save period in seconds

	*/
	Period int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service row cache save period post params
func (o *CacheServiceRowCacheSavePeriodPostParams) WithTimeout(timeout time.Duration) *CacheServiceRowCacheSavePeriodPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service row cache save period post params
func (o *CacheServiceRowCacheSavePeriodPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service row cache save period post params
func (o *CacheServiceRowCacheSavePeriodPostParams) WithContext(ctx context.Context) *CacheServiceRowCacheSavePeriodPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service row cache save period post params
func (o *CacheServiceRowCacheSavePeriodPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service row cache save period post params
func (o *CacheServiceRowCacheSavePeriodPostParams) WithHTTPClient(client *http.Client) *CacheServiceRowCacheSavePeriodPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service row cache save period post params
func (o *CacheServiceRowCacheSavePeriodPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPeriod adds the period to the cache service row cache save period post params
func (o *CacheServiceRowCacheSavePeriodPostParams) WithPeriod(period int32) *CacheServiceRowCacheSavePeriodPostParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the cache service row cache save period post params
func (o *CacheServiceRowCacheSavePeriodPostParams) SetPeriod(period int32) {
	o.Period = period
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceRowCacheSavePeriodPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param period
	qrPeriod := o.Period
	qPeriod := swag.FormatInt32(qrPeriod)
	if qPeriod != "" {
		if err := r.SetQueryParam("period", qPeriod); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
