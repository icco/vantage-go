// Code generated by go-swagger; DO NOT EDIT.

package prices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// GetProductsReader is a Reader for the GetProducts structure.
type GetProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /products] getProducts", response, response.Code())
	}
}

// NewGetProductsOK creates a GetProductsOK with default headers values
func NewGetProductsOK() *GetProductsOK {
	return &GetProductsOK{}
}

/*
GetProductsOK describes a response with status code 200, with default header values.

GetProductsOK get products o k
*/
type GetProductsOK struct {
	Payload *models.Products
}

// IsSuccess returns true when this get products o k response has a 2xx status code
func (o *GetProductsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get products o k response has a 3xx status code
func (o *GetProductsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get products o k response has a 4xx status code
func (o *GetProductsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get products o k response has a 5xx status code
func (o *GetProductsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get products o k response a status code equal to that given
func (o *GetProductsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get products o k response
func (o *GetProductsOK) Code() int {
	return 200
}

func (o *GetProductsOK) Error() string {
	return fmt.Sprintf("[GET /products][%d] getProductsOK  %+v", 200, o.Payload)
}

func (o *GetProductsOK) String() string {
	return fmt.Sprintf("[GET /products][%d] getProductsOK  %+v", 200, o.Payload)
}

func (o *GetProductsOK) GetPayload() *models.Products {
	return o.Payload
}

func (o *GetProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Products)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
