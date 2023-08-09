// Code generated by go-swagger; DO NOT EDIT.

package costs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev1/models"
)

// GetReportReader is a Reader for the GetReport structure.
type GetReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /reports/{report_id}] getReport", response, response.Code())
	}
}

// NewGetReportOK creates a GetReportOK with default headers values
func NewGetReportOK() *GetReportOK {
	return &GetReportOK{}
}

/*
GetReportOK describes a response with status code 200, with default header values.

GetReportOK get report o k
*/
type GetReportOK struct {
	Payload *models.Report
}

// IsSuccess returns true when this get report o k response has a 2xx status code
func (o *GetReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get report o k response has a 3xx status code
func (o *GetReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report o k response has a 4xx status code
func (o *GetReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get report o k response has a 5xx status code
func (o *GetReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get report o k response a status code equal to that given
func (o *GetReportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get report o k response
func (o *GetReportOK) Code() int {
	return 200
}

func (o *GetReportOK) Error() string {
	return fmt.Sprintf("[GET /reports/{report_id}][%d] getReportOK  %+v", 200, o.Payload)
}

func (o *GetReportOK) String() string {
	return fmt.Sprintf("[GET /reports/{report_id}][%d] getReportOK  %+v", 200, o.Payload)
}

func (o *GetReportOK) GetPayload() *models.Report {
	return o.Payload
}

func (o *GetReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Report)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
