// Code generated by go-swagger; DO NOT EDIT.

package constraint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListDefaultConstraintReader is a Reader for the ListDefaultConstraint structure.
type ListDefaultConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDefaultConstraintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDefaultConstraintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListDefaultConstraintUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListDefaultConstraintForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListDefaultConstraintDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDefaultConstraintOK creates a ListDefaultConstraintOK with default headers values
func NewListDefaultConstraintOK() *ListDefaultConstraintOK {
	return &ListDefaultConstraintOK{}
}

/* ListDefaultConstraintOK describes a response with status code 200, with default header values.

Constraint
*/
type ListDefaultConstraintOK struct {
	Payload []*models.Constraint
}

func (o *ListDefaultConstraintOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/constraints][%d] listDefaultConstraintOK  %+v", 200, o.Payload)
}
func (o *ListDefaultConstraintOK) GetPayload() []*models.Constraint {
	return o.Payload
}

func (o *ListDefaultConstraintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDefaultConstraintUnauthorized creates a ListDefaultConstraintUnauthorized with default headers values
func NewListDefaultConstraintUnauthorized() *ListDefaultConstraintUnauthorized {
	return &ListDefaultConstraintUnauthorized{}
}

/* ListDefaultConstraintUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListDefaultConstraintUnauthorized struct {
}

func (o *ListDefaultConstraintUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/constraints][%d] listDefaultConstraintUnauthorized ", 401)
}

func (o *ListDefaultConstraintUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListDefaultConstraintForbidden creates a ListDefaultConstraintForbidden with default headers values
func NewListDefaultConstraintForbidden() *ListDefaultConstraintForbidden {
	return &ListDefaultConstraintForbidden{}
}

/* ListDefaultConstraintForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListDefaultConstraintForbidden struct {
}

func (o *ListDefaultConstraintForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/constraints][%d] listDefaultConstraintForbidden ", 403)
}

func (o *ListDefaultConstraintForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListDefaultConstraintDefault creates a ListDefaultConstraintDefault with default headers values
func NewListDefaultConstraintDefault(code int) *ListDefaultConstraintDefault {
	return &ListDefaultConstraintDefault{
		_statusCode: code,
	}
}

/* ListDefaultConstraintDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListDefaultConstraintDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list default constraint default response
func (o *ListDefaultConstraintDefault) Code() int {
	return o._statusCode
}

func (o *ListDefaultConstraintDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/constraints][%d] listDefaultConstraint default  %+v", o._statusCode, o.Payload)
}
func (o *ListDefaultConstraintDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListDefaultConstraintDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
