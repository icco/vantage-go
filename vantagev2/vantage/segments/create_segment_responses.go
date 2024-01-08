// Code generated by go-swagger; DO NOT EDIT.

package segments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// CreateSegmentReader is a Reader for the CreateSegment structure.
type CreateSegmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSegmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSegmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSegmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateSegmentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /segments] createSegment", response, response.Code())
	}
}

// NewCreateSegmentCreated creates a CreateSegmentCreated with default headers values
func NewCreateSegmentCreated() *CreateSegmentCreated {
	return &CreateSegmentCreated{}
}

/*
CreateSegmentCreated describes a response with status code 201, with default header values.

CreateSegmentCreated create segment created
*/
type CreateSegmentCreated struct {
	Payload *models.Segment
}

// IsSuccess returns true when this create segment created response has a 2xx status code
func (o *CreateSegmentCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create segment created response has a 3xx status code
func (o *CreateSegmentCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create segment created response has a 4xx status code
func (o *CreateSegmentCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create segment created response has a 5xx status code
func (o *CreateSegmentCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create segment created response a status code equal to that given
func (o *CreateSegmentCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create segment created response
func (o *CreateSegmentCreated) Code() int {
	return 201
}

func (o *CreateSegmentCreated) Error() string {
	return fmt.Sprintf("[POST /segments][%d] createSegmentCreated  %+v", 201, o.Payload)
}

func (o *CreateSegmentCreated) String() string {
	return fmt.Sprintf("[POST /segments][%d] createSegmentCreated  %+v", 201, o.Payload)
}

func (o *CreateSegmentCreated) GetPayload() *models.Segment {
	return o.Payload
}

func (o *CreateSegmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Segment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSegmentBadRequest creates a CreateSegmentBadRequest with default headers values
func NewCreateSegmentBadRequest() *CreateSegmentBadRequest {
	return &CreateSegmentBadRequest{}
}

/*
CreateSegmentBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type CreateSegmentBadRequest struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create segment bad request response has a 2xx status code
func (o *CreateSegmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create segment bad request response has a 3xx status code
func (o *CreateSegmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create segment bad request response has a 4xx status code
func (o *CreateSegmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create segment bad request response has a 5xx status code
func (o *CreateSegmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create segment bad request response a status code equal to that given
func (o *CreateSegmentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create segment bad request response
func (o *CreateSegmentBadRequest) Code() int {
	return 400
}

func (o *CreateSegmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /segments][%d] createSegmentBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSegmentBadRequest) String() string {
	return fmt.Sprintf("[POST /segments][%d] createSegmentBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSegmentBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateSegmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSegmentUnprocessableEntity creates a CreateSegmentUnprocessableEntity with default headers values
func NewCreateSegmentUnprocessableEntity() *CreateSegmentUnprocessableEntity {
	return &CreateSegmentUnprocessableEntity{}
}

/*
CreateSegmentUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntity
*/
type CreateSegmentUnprocessableEntity struct {
	Payload *models.Errors
}

// IsSuccess returns true when this create segment unprocessable entity response has a 2xx status code
func (o *CreateSegmentUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create segment unprocessable entity response has a 3xx status code
func (o *CreateSegmentUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create segment unprocessable entity response has a 4xx status code
func (o *CreateSegmentUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create segment unprocessable entity response has a 5xx status code
func (o *CreateSegmentUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create segment unprocessable entity response a status code equal to that given
func (o *CreateSegmentUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create segment unprocessable entity response
func (o *CreateSegmentUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateSegmentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /segments][%d] createSegmentUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateSegmentUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /segments][%d] createSegmentUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateSegmentUnprocessableEntity) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateSegmentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
