// Code generated by go-swagger; DO NOT EDIT.

package financial_commitments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

// GetFinancialCommitmentsReader is a Reader for the GetFinancialCommitments structure.
type GetFinancialCommitmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFinancialCommitmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFinancialCommitmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /financial_commitments] getFinancialCommitments", response, response.Code())
	}
}

// NewGetFinancialCommitmentsOK creates a GetFinancialCommitmentsOK with default headers values
func NewGetFinancialCommitmentsOK() *GetFinancialCommitmentsOK {
	return &GetFinancialCommitmentsOK{}
}

/*
GetFinancialCommitmentsOK describes a response with status code 200, with default header values.

GetFinancialCommitmentsOK get financial commitments o k
*/
type GetFinancialCommitmentsOK struct {
	Payload *models.FinancialCommitments
}

// IsSuccess returns true when this get financial commitments o k response has a 2xx status code
func (o *GetFinancialCommitmentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get financial commitments o k response has a 3xx status code
func (o *GetFinancialCommitmentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get financial commitments o k response has a 4xx status code
func (o *GetFinancialCommitmentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get financial commitments o k response has a 5xx status code
func (o *GetFinancialCommitmentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get financial commitments o k response a status code equal to that given
func (o *GetFinancialCommitmentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get financial commitments o k response
func (o *GetFinancialCommitmentsOK) Code() int {
	return 200
}

func (o *GetFinancialCommitmentsOK) Error() string {
	return fmt.Sprintf("[GET /financial_commitments][%d] getFinancialCommitmentsOK  %+v", 200, o.Payload)
}

func (o *GetFinancialCommitmentsOK) String() string {
	return fmt.Sprintf("[GET /financial_commitments][%d] getFinancialCommitmentsOK  %+v", 200, o.Payload)
}

func (o *GetFinancialCommitmentsOK) GetPayload() *models.FinancialCommitments {
	return o.Payload
}

func (o *GetFinancialCommitmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FinancialCommitments)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
