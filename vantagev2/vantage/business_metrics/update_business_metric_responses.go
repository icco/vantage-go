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

// UpdateBusinessMetricReader is a Reader for the UpdateBusinessMetric structure.
type UpdateBusinessMetricReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBusinessMetricReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateBusinessMetricCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateBusinessMetricBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateBusinessMetricForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateBusinessMetricNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateBusinessMetricUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /business_metrics/{business_metric_token}] updateBusinessMetric", response, response.Code())
	}
}

// NewUpdateBusinessMetricCreated creates a UpdateBusinessMetricCreated with default headers values
func NewUpdateBusinessMetricCreated() *UpdateBusinessMetricCreated {
	return &UpdateBusinessMetricCreated{}
}

/*
UpdateBusinessMetricCreated describes a response with status code 201, with default header values.

UpdateBusinessMetricCreated update business metric created
*/
type UpdateBusinessMetricCreated struct {
	Payload *models.BusinessMetric
}

// IsSuccess returns true when this update business metric created response has a 2xx status code
func (o *UpdateBusinessMetricCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update business metric created response has a 3xx status code
func (o *UpdateBusinessMetricCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update business metric created response has a 4xx status code
func (o *UpdateBusinessMetricCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update business metric created response has a 5xx status code
func (o *UpdateBusinessMetricCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update business metric created response a status code equal to that given
func (o *UpdateBusinessMetricCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update business metric created response
func (o *UpdateBusinessMetricCreated) Code() int {
	return 201
}

func (o *UpdateBusinessMetricCreated) Error() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}][%d] updateBusinessMetricCreated  %+v", 201, o.Payload)
}

func (o *UpdateBusinessMetricCreated) String() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}][%d] updateBusinessMetricCreated  %+v", 201, o.Payload)
}

func (o *UpdateBusinessMetricCreated) GetPayload() *models.BusinessMetric {
	return o.Payload
}

func (o *UpdateBusinessMetricCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BusinessMetric)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBusinessMetricBadRequest creates a UpdateBusinessMetricBadRequest with default headers values
func NewUpdateBusinessMetricBadRequest() *UpdateBusinessMetricBadRequest {
	return &UpdateBusinessMetricBadRequest{}
}

/*
UpdateBusinessMetricBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateBusinessMetricBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update business metric bad request response has a 2xx status code
func (o *UpdateBusinessMetricBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update business metric bad request response has a 3xx status code
func (o *UpdateBusinessMetricBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update business metric bad request response has a 4xx status code
func (o *UpdateBusinessMetricBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update business metric bad request response has a 5xx status code
func (o *UpdateBusinessMetricBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update business metric bad request response a status code equal to that given
func (o *UpdateBusinessMetricBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update business metric bad request response
func (o *UpdateBusinessMetricBadRequest) Code() int {
	return 400
}

func (o *UpdateBusinessMetricBadRequest) Error() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}][%d] updateBusinessMetricBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateBusinessMetricBadRequest) String() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}][%d] updateBusinessMetricBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateBusinessMetricBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateBusinessMetricBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBusinessMetricForbidden creates a UpdateBusinessMetricForbidden with default headers values
func NewUpdateBusinessMetricForbidden() *UpdateBusinessMetricForbidden {
	return &UpdateBusinessMetricForbidden{}
}

/*
UpdateBusinessMetricForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateBusinessMetricForbidden struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update business metric forbidden response has a 2xx status code
func (o *UpdateBusinessMetricForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update business metric forbidden response has a 3xx status code
func (o *UpdateBusinessMetricForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update business metric forbidden response has a 4xx status code
func (o *UpdateBusinessMetricForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update business metric forbidden response has a 5xx status code
func (o *UpdateBusinessMetricForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update business metric forbidden response a status code equal to that given
func (o *UpdateBusinessMetricForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update business metric forbidden response
func (o *UpdateBusinessMetricForbidden) Code() int {
	return 403
}

func (o *UpdateBusinessMetricForbidden) Error() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}][%d] updateBusinessMetricForbidden  %+v", 403, o.Payload)
}

func (o *UpdateBusinessMetricForbidden) String() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}][%d] updateBusinessMetricForbidden  %+v", 403, o.Payload)
}

func (o *UpdateBusinessMetricForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateBusinessMetricForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBusinessMetricNotFound creates a UpdateBusinessMetricNotFound with default headers values
func NewUpdateBusinessMetricNotFound() *UpdateBusinessMetricNotFound {
	return &UpdateBusinessMetricNotFound{}
}

/*
UpdateBusinessMetricNotFound describes a response with status code 404, with default header values.

NotFound
*/
type UpdateBusinessMetricNotFound struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update business metric not found response has a 2xx status code
func (o *UpdateBusinessMetricNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update business metric not found response has a 3xx status code
func (o *UpdateBusinessMetricNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update business metric not found response has a 4xx status code
func (o *UpdateBusinessMetricNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update business metric not found response has a 5xx status code
func (o *UpdateBusinessMetricNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update business metric not found response a status code equal to that given
func (o *UpdateBusinessMetricNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update business metric not found response
func (o *UpdateBusinessMetricNotFound) Code() int {
	return 404
}

func (o *UpdateBusinessMetricNotFound) Error() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}][%d] updateBusinessMetricNotFound  %+v", 404, o.Payload)
}

func (o *UpdateBusinessMetricNotFound) String() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}][%d] updateBusinessMetricNotFound  %+v", 404, o.Payload)
}

func (o *UpdateBusinessMetricNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateBusinessMetricNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBusinessMetricUnprocessableEntity creates a UpdateBusinessMetricUnprocessableEntity with default headers values
func NewUpdateBusinessMetricUnprocessableEntity() *UpdateBusinessMetricUnprocessableEntity {
	return &UpdateBusinessMetricUnprocessableEntity{}
}

/*
UpdateBusinessMetricUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntity
*/
type UpdateBusinessMetricUnprocessableEntity struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update business metric unprocessable entity response has a 2xx status code
func (o *UpdateBusinessMetricUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update business metric unprocessable entity response has a 3xx status code
func (o *UpdateBusinessMetricUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update business metric unprocessable entity response has a 4xx status code
func (o *UpdateBusinessMetricUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update business metric unprocessable entity response has a 5xx status code
func (o *UpdateBusinessMetricUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update business metric unprocessable entity response a status code equal to that given
func (o *UpdateBusinessMetricUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update business metric unprocessable entity response
func (o *UpdateBusinessMetricUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateBusinessMetricUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}][%d] updateBusinessMetricUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateBusinessMetricUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /business_metrics/{business_metric_token}][%d] updateBusinessMetricUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateBusinessMetricUnprocessableEntity) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateBusinessMetricUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
