// Code generated by go-swagger; DO NOT EDIT.

package prices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev1/models"
)

// GetProductReader is a Reader for the GetProduct structure.
type GetProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /products/{id}] getProduct", response, response.Code())
	}
}

// NewGetProductOK creates a GetProductOK with default headers values
func NewGetProductOK() *GetProductOK {
	return &GetProductOK{}
}

/*
GetProductOK describes a response with status code 200, with default header values.

GetProductOK get product o k
*/
type GetProductOK struct {
	Payload *models.Product
}

// IsSuccess returns true when this get product o k response has a 2xx status code
func (o *GetProductOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get product o k response has a 3xx status code
func (o *GetProductOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get product o k response has a 4xx status code
func (o *GetProductOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get product o k response has a 5xx status code
func (o *GetProductOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get product o k response a status code equal to that given
func (o *GetProductOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get product o k response
func (o *GetProductOK) Code() int {
	return 200
}

func (o *GetProductOK) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductOK  %+v", 200, o.Payload)
}

func (o *GetProductOK) String() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductOK  %+v", 200, o.Payload)
}

func (o *GetProductOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *GetProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
