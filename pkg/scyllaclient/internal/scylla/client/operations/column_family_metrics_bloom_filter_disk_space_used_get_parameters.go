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

// NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams creates a new ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams() *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams {

	return &ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParamsWithTimeout creates a new ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams {

	return &ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParamsWithContext creates a new ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams {

	return &ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParamsWithHTTPClient creates a new ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams {

	return &ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams contains all the parameters to send to the API endpoint
for the column family metrics bloom filter disk space used get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics bloom filter disk space used get params
func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics bloom filter disk space used get params
func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics bloom filter disk space used get params
func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics bloom filter disk space used get params
func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics bloom filter disk space used get params
func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics bloom filter disk space used get params
func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
