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

// StorageServiceTokensGetReader is a Reader for the StorageServiceTokensGet structure.
type StorageServiceTokensGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceTokensGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceTokensGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceTokensGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceTokensGetOK creates a StorageServiceTokensGetOK with default headers values
func NewStorageServiceTokensGetOK() *StorageServiceTokensGetOK {
	return &StorageServiceTokensGetOK{}
}

/*
StorageServiceTokensGetOK handles this case with default header values.

StorageServiceTokensGetOK storage service tokens get o k
*/
type StorageServiceTokensGetOK struct {
	Payload []string
}

func (o *StorageServiceTokensGetOK) GetPayload() []string {
	return o.Payload
}

func (o *StorageServiceTokensGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceTokensGetDefault creates a StorageServiceTokensGetDefault with default headers values
func NewStorageServiceTokensGetDefault(code int) *StorageServiceTokensGetDefault {
	return &StorageServiceTokensGetDefault{
		_statusCode: code,
	}
}

/*
StorageServiceTokensGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceTokensGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service tokens get default response
func (o *StorageServiceTokensGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceTokensGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceTokensGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceTokensGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
