// Code generated by go-swagger; DO NOT EDIT.

package costs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// GetDashboardsReader is a Reader for the GetDashboards structure.
type GetDashboardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /dashboards] getDashboards", response, response.Code())
	}
}

// NewGetDashboardsOK creates a GetDashboardsOK with default headers values
func NewGetDashboardsOK() *GetDashboardsOK {
	return &GetDashboardsOK{}
}

/*
GetDashboardsOK describes a response with status code 200, with default header values.

GetDashboardsOK get dashboards o k
*/
type GetDashboardsOK struct {
	Payload *models.Dashboards
}

// IsSuccess returns true when this get dashboards o k response has a 2xx status code
func (o *GetDashboardsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dashboards o k response has a 3xx status code
func (o *GetDashboardsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboards o k response has a 4xx status code
func (o *GetDashboardsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dashboards o k response has a 5xx status code
func (o *GetDashboardsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboards o k response a status code equal to that given
func (o *GetDashboardsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get dashboards o k response
func (o *GetDashboardsOK) Code() int {
	return 200
}

func (o *GetDashboardsOK) Error() string {
	return fmt.Sprintf("[GET /dashboards][%d] getDashboardsOK  %+v", 200, o.Payload)
}

func (o *GetDashboardsOK) String() string {
	return fmt.Sprintf("[GET /dashboards][%d] getDashboardsOK  %+v", 200, o.Payload)
}

func (o *GetDashboardsOK) GetPayload() *models.Dashboards {
	return o.Payload
}

func (o *GetDashboardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Dashboards)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
