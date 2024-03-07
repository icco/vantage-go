// Code generated by go-swagger; DO NOT EDIT.

package business_metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// CreateBusinessMetricReader is a Reader for the CreateBusinessMetric structure.
type CreateBusinessMetricReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBusinessMetricReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateBusinessMetricCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateBusinessMetricBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateBusinessMetricForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateBusinessMetricNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateBusinessMetricUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /business_metrics] createBusinessMetric", response, response.Code())
	}
}

// NewCreateBusinessMetricCreated creates a CreateBusinessMetricCreated with default headers values
func NewCreateBusinessMetricCreated() *CreateBusinessMetricCreated {
	return &CreateBusinessMetricCreated{}
}

/*
CreateBusinessMetricCreated describes a response with status code 201, with default header values.

CreateBusinessMetricCreated create business metric created
*/
type CreateBusinessMetricCreated struct {
	Payload *models.BusinessMetric
}

// IsSuccess returns true when this create business metric created response has a 2xx status code
func (o *CreateBusinessMetricCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create business metric created response has a 3xx status code
func (o *CreateBusinessMetricCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create business metric created response has a 4xx status code
func (o *CreateBusinessMetricCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create business metric created response has a 5xx status code
func (o *CreateBusinessMetricCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create business metric created response a status code equal to that given
func (o *CreateBusinessMetricCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create business metric created response
func (o *CreateBusinessMetricCreated) Code() int {
	return 201
}

func (o *CreateBusinessMetricCreated) Error() string {
	return fmt.Sprintf("[POST /business_metrics][%d] createBusinessMetricCreated  %+v", 201, o.Payload)
}

func (o *CreateBusinessMetricCreated) String() string {
	return fmt.Sprintf("[POST /business_metrics][%d] createBusinessMetricCreated  %+v", 201, o.Payload)
}

func (o *CreateBusinessMetricCreated) GetPayload() *models.BusinessMetric {
	return o.Payload
}

func (o *CreateBusinessMetricCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BusinessMetric)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBusinessMetricBadRequest creates a CreateBusinessMetricBadRequest with default headers values
func NewCreateBusinessMetricBadRequest() *CreateBusinessMetricBadRequest {
	return &CreateBusinessMetricBadRequest{}
}

/*
CreateBusinessMetricBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type CreateBusinessMetricBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create business metric bad request response has a 2xx status code
func (o *CreateBusinessMetricBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create business metric bad request response has a 3xx status code
func (o *CreateBusinessMetricBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create business metric bad request response has a 4xx status code
func (o *CreateBusinessMetricBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create business metric bad request response has a 5xx status code
func (o *CreateBusinessMetricBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create business metric bad request response a status code equal to that given
func (o *CreateBusinessMetricBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create business metric bad request response
func (o *CreateBusinessMetricBadRequest) Code() int {
	return 400
}

func (o *CreateBusinessMetricBadRequest) Error() string {
	return fmt.Sprintf("[POST /business_metrics][%d] createBusinessMetricBadRequest  %+v", 400, o.Payload)
}

func (o *CreateBusinessMetricBadRequest) String() string {
	return fmt.Sprintf("[POST /business_metrics][%d] createBusinessMetricBadRequest  %+v", 400, o.Payload)
}

func (o *CreateBusinessMetricBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateBusinessMetricBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBusinessMetricForbidden creates a CreateBusinessMetricForbidden with default headers values
func NewCreateBusinessMetricForbidden() *CreateBusinessMetricForbidden {
	return &CreateBusinessMetricForbidden{}
}

/*
CreateBusinessMetricForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateBusinessMetricForbidden struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create business metric forbidden response has a 2xx status code
func (o *CreateBusinessMetricForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create business metric forbidden response has a 3xx status code
func (o *CreateBusinessMetricForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create business metric forbidden response has a 4xx status code
func (o *CreateBusinessMetricForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create business metric forbidden response has a 5xx status code
func (o *CreateBusinessMetricForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create business metric forbidden response a status code equal to that given
func (o *CreateBusinessMetricForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create business metric forbidden response
func (o *CreateBusinessMetricForbidden) Code() int {
	return 403
}

func (o *CreateBusinessMetricForbidden) Error() string {
	return fmt.Sprintf("[POST /business_metrics][%d] createBusinessMetricForbidden  %+v", 403, o.Payload)
}

func (o *CreateBusinessMetricForbidden) String() string {
	return fmt.Sprintf("[POST /business_metrics][%d] createBusinessMetricForbidden  %+v", 403, o.Payload)
}

func (o *CreateBusinessMetricForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateBusinessMetricForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBusinessMetricNotFound creates a CreateBusinessMetricNotFound with default headers values
func NewCreateBusinessMetricNotFound() *CreateBusinessMetricNotFound {
	return &CreateBusinessMetricNotFound{}
}

/*
CreateBusinessMetricNotFound describes a response with status code 404, with default header values.

NotFound
*/
type CreateBusinessMetricNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create business metric not found response has a 2xx status code
func (o *CreateBusinessMetricNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create business metric not found response has a 3xx status code
func (o *CreateBusinessMetricNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create business metric not found response has a 4xx status code
func (o *CreateBusinessMetricNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create business metric not found response has a 5xx status code
func (o *CreateBusinessMetricNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create business metric not found response a status code equal to that given
func (o *CreateBusinessMetricNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create business metric not found response
func (o *CreateBusinessMetricNotFound) Code() int {
	return 404
}

func (o *CreateBusinessMetricNotFound) Error() string {
	return fmt.Sprintf("[POST /business_metrics][%d] createBusinessMetricNotFound  %+v", 404, o.Payload)
}

func (o *CreateBusinessMetricNotFound) String() string {
	return fmt.Sprintf("[POST /business_metrics][%d] createBusinessMetricNotFound  %+v", 404, o.Payload)
}

func (o *CreateBusinessMetricNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateBusinessMetricNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBusinessMetricUnprocessableEntity creates a CreateBusinessMetricUnprocessableEntity with default headers values
func NewCreateBusinessMetricUnprocessableEntity() *CreateBusinessMetricUnprocessableEntity {
	return &CreateBusinessMetricUnprocessableEntity{}
}

/*
CreateBusinessMetricUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntity
*/
type CreateBusinessMetricUnprocessableEntity struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create business metric unprocessable entity response has a 2xx status code
func (o *CreateBusinessMetricUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create business metric unprocessable entity response has a 3xx status code
func (o *CreateBusinessMetricUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create business metric unprocessable entity response has a 4xx status code
func (o *CreateBusinessMetricUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create business metric unprocessable entity response has a 5xx status code
func (o *CreateBusinessMetricUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create business metric unprocessable entity response a status code equal to that given
func (o *CreateBusinessMetricUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create business metric unprocessable entity response
func (o *CreateBusinessMetricUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateBusinessMetricUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /business_metrics][%d] createBusinessMetricUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateBusinessMetricUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /business_metrics][%d] createBusinessMetricUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateBusinessMetricUnprocessableEntity) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateBusinessMetricUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
