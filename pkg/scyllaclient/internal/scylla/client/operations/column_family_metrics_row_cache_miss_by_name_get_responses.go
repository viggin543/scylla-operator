// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// ColumnFamilyMetricsRowCacheMissByNameGetReader is a Reader for the ColumnFamilyMetricsRowCacheMissByNameGet structure.
type ColumnFamilyMetricsRowCacheMissByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsRowCacheMissByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsRowCacheMissByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsRowCacheMissByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsRowCacheMissByNameGetOK creates a ColumnFamilyMetricsRowCacheMissByNameGetOK with default headers values
func NewColumnFamilyMetricsRowCacheMissByNameGetOK() *ColumnFamilyMetricsRowCacheMissByNameGetOK {
	return &ColumnFamilyMetricsRowCacheMissByNameGetOK{}
}

/*
ColumnFamilyMetricsRowCacheMissByNameGetOK handles this case with default header values.

ColumnFamilyMetricsRowCacheMissByNameGetOK column family metrics row cache miss by name get o k
*/
type ColumnFamilyMetricsRowCacheMissByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsRowCacheMissByNameGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsRowCacheMissByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsRowCacheMissByNameGetDefault creates a ColumnFamilyMetricsRowCacheMissByNameGetDefault with default headers values
func NewColumnFamilyMetricsRowCacheMissByNameGetDefault(code int) *ColumnFamilyMetricsRowCacheMissByNameGetDefault {
	return &ColumnFamilyMetricsRowCacheMissByNameGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsRowCacheMissByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsRowCacheMissByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics row cache miss by name get default response
func (o *ColumnFamilyMetricsRowCacheMissByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsRowCacheMissByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsRowCacheMissByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsRowCacheMissByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
