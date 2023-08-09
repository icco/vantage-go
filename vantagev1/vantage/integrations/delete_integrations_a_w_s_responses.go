// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteIntegrationsAWSReader is a Reader for the DeleteIntegrationsAWS structure.
type DeleteIntegrationsAWSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIntegrationsAWSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteIntegrationsAWSNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /integrations/aws/{access_credential_id}] deleteIntegrationsAWS", response, response.Code())
	}
}

// NewDeleteIntegrationsAWSNoContent creates a DeleteIntegrationsAWSNoContent with default headers values
func NewDeleteIntegrationsAWSNoContent() *DeleteIntegrationsAWSNoContent {
	return &DeleteIntegrationsAWSNoContent{}
}

/*
DeleteIntegrationsAWSNoContent describes a response with status code 204, with default header values.

DeleteIntegrationsAWSNoContent delete integrations a w s no content
*/
type DeleteIntegrationsAWSNoContent struct {
}

// IsSuccess returns true when this delete integrations a w s no content response has a 2xx status code
func (o *DeleteIntegrationsAWSNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete integrations a w s no content response has a 3xx status code
func (o *DeleteIntegrationsAWSNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integrations a w s no content response has a 4xx status code
func (o *DeleteIntegrationsAWSNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete integrations a w s no content response has a 5xx status code
func (o *DeleteIntegrationsAWSNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integrations a w s no content response a status code equal to that given
func (o *DeleteIntegrationsAWSNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete integrations a w s no content response
func (o *DeleteIntegrationsAWSNoContent) Code() int {
	return 204
}

func (o *DeleteIntegrationsAWSNoContent) Error() string {
	return fmt.Sprintf("[DELETE /integrations/aws/{access_credential_id}][%d] deleteIntegrationsAWSNoContent ", 204)
}

func (o *DeleteIntegrationsAWSNoContent) String() string {
	return fmt.Sprintf("[DELETE /integrations/aws/{access_credential_id}][%d] deleteIntegrationsAWSNoContent ", 204)
}

func (o *DeleteIntegrationsAWSNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
