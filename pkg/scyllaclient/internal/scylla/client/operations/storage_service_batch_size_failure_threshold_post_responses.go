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

// StorageServiceBatchSizeFailureThresholdPostReader is a Reader for the StorageServiceBatchSizeFailureThresholdPost structure.
type StorageServiceBatchSizeFailureThresholdPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceBatchSizeFailureThresholdPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceBatchSizeFailureThresholdPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceBatchSizeFailureThresholdPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceBatchSizeFailureThresholdPostOK creates a StorageServiceBatchSizeFailureThresholdPostOK with default headers values
func NewStorageServiceBatchSizeFailureThresholdPostOK() *StorageServiceBatchSizeFailureThresholdPostOK {
	return &StorageServiceBatchSizeFailureThresholdPostOK{}
}

/*
StorageServiceBatchSizeFailureThresholdPostOK handles this case with default header values.

StorageServiceBatchSizeFailureThresholdPostOK storage service batch size failure threshold post o k
*/
type StorageServiceBatchSizeFailureThresholdPostOK struct {
}

func (o *StorageServiceBatchSizeFailureThresholdPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceBatchSizeFailureThresholdPostDefault creates a StorageServiceBatchSizeFailureThresholdPostDefault with default headers values
func NewStorageServiceBatchSizeFailureThresholdPostDefault(code int) *StorageServiceBatchSizeFailureThresholdPostDefault {
	return &StorageServiceBatchSizeFailureThresholdPostDefault{
		_statusCode: code,
	}
}

/*
StorageServiceBatchSizeFailureThresholdPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceBatchSizeFailureThresholdPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service batch size failure threshold post default response
func (o *StorageServiceBatchSizeFailureThresholdPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceBatchSizeFailureThresholdPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceBatchSizeFailureThresholdPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceBatchSizeFailureThresholdPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
