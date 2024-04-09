// Code generated by go-swagger; DO NOT EDIT.

package anomaly_notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// UpdateAnomalyNotificationReader is a Reader for the UpdateAnomalyNotification structure.
type UpdateAnomalyNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAnomalyNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAnomalyNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAnomalyNotificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /anomaly_notifications/{anomaly_notification_token}] updateAnomalyNotification", response, response.Code())
	}
}

// NewUpdateAnomalyNotificationOK creates a UpdateAnomalyNotificationOK with default headers values
func NewUpdateAnomalyNotificationOK() *UpdateAnomalyNotificationOK {
	return &UpdateAnomalyNotificationOK{}
}

/*
UpdateAnomalyNotificationOK describes a response with status code 200, with default header values.

UpdateAnomalyNotificationOK update anomaly notification o k
*/
type UpdateAnomalyNotificationOK struct {
	Payload *models.AnomalyNotification
}

// IsSuccess returns true when this update anomaly notification o k response has a 2xx status code
func (o *UpdateAnomalyNotificationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update anomaly notification o k response has a 3xx status code
func (o *UpdateAnomalyNotificationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update anomaly notification o k response has a 4xx status code
func (o *UpdateAnomalyNotificationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update anomaly notification o k response has a 5xx status code
func (o *UpdateAnomalyNotificationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update anomaly notification o k response a status code equal to that given
func (o *UpdateAnomalyNotificationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update anomaly notification o k response
func (o *UpdateAnomalyNotificationOK) Code() int {
	return 200
}

func (o *UpdateAnomalyNotificationOK) Error() string {
	return fmt.Sprintf("[PUT /anomaly_notifications/{anomaly_notification_token}][%d] updateAnomalyNotificationOK  %+v", 200, o.Payload)
}

func (o *UpdateAnomalyNotificationOK) String() string {
	return fmt.Sprintf("[PUT /anomaly_notifications/{anomaly_notification_token}][%d] updateAnomalyNotificationOK  %+v", 200, o.Payload)
}

func (o *UpdateAnomalyNotificationOK) GetPayload() *models.AnomalyNotification {
	return o.Payload
}

func (o *UpdateAnomalyNotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnomalyNotification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAnomalyNotificationBadRequest creates a UpdateAnomalyNotificationBadRequest with default headers values
func NewUpdateAnomalyNotificationBadRequest() *UpdateAnomalyNotificationBadRequest {
	return &UpdateAnomalyNotificationBadRequest{}
}

/*
UpdateAnomalyNotificationBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateAnomalyNotificationBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this update anomaly notification bad request response has a 2xx status code
func (o *UpdateAnomalyNotificationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update anomaly notification bad request response has a 3xx status code
func (o *UpdateAnomalyNotificationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update anomaly notification bad request response has a 4xx status code
func (o *UpdateAnomalyNotificationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update anomaly notification bad request response has a 5xx status code
func (o *UpdateAnomalyNotificationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update anomaly notification bad request response a status code equal to that given
func (o *UpdateAnomalyNotificationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update anomaly notification bad request response
func (o *UpdateAnomalyNotificationBadRequest) Code() int {
	return 400
}

func (o *UpdateAnomalyNotificationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /anomaly_notifications/{anomaly_notification_token}][%d] updateAnomalyNotificationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAnomalyNotificationBadRequest) String() string {
	return fmt.Sprintf("[PUT /anomaly_notifications/{anomaly_notification_token}][%d] updateAnomalyNotificationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAnomalyNotificationBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateAnomalyNotificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
