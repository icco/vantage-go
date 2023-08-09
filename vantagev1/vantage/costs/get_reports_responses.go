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

// GetReportsReader is a Reader for the GetReports structure.
type GetReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /reports] getReports", response, response.Code())
	}
}

// NewGetReportsOK creates a GetReportsOK with default headers values
func NewGetReportsOK() *GetReportsOK {
	return &GetReportsOK{}
}

/*
GetReportsOK describes a response with status code 200, with default header values.

GetReportsOK get reports o k
*/
type GetReportsOK struct {
	Payload *models.Reports
}

// IsSuccess returns true when this get reports o k response has a 2xx status code
func (o *GetReportsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get reports o k response has a 3xx status code
func (o *GetReportsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get reports o k response has a 4xx status code
func (o *GetReportsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get reports o k response has a 5xx status code
func (o *GetReportsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get reports o k response a status code equal to that given
func (o *GetReportsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get reports o k response
func (o *GetReportsOK) Code() int {
	return 200
}

func (o *GetReportsOK) Error() string {
	return fmt.Sprintf("[GET /reports][%d] getReportsOK  %+v", 200, o.Payload)
}

func (o *GetReportsOK) String() string {
	return fmt.Sprintf("[GET /reports][%d] getReportsOK  %+v", 200, o.Payload)
}

func (o *GetReportsOK) GetPayload() *models.Reports {
	return o.Payload
}

func (o *GetReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Reports)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
