// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kata-containers/kata-containers/src/runtime/virtcontainers/pkg/firecracker/client/models"
)

// PatchGuestDriveByIDReader is a Reader for the PatchGuestDriveByID structure.
type PatchGuestDriveByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchGuestDriveByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchGuestDriveByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchGuestDriveByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchGuestDriveByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchGuestDriveByIDNoContent creates a PatchGuestDriveByIDNoContent with default headers values
func NewPatchGuestDriveByIDNoContent() *PatchGuestDriveByIDNoContent {
	return &PatchGuestDriveByIDNoContent{}
}

/*PatchGuestDriveByIDNoContent handles this case with default header values.

Drive updated
*/
type PatchGuestDriveByIDNoContent struct {
}

func (o *PatchGuestDriveByIDNoContent) Error() string {
	return fmt.Sprintf("[PATCH /drives/{drive_id}][%d] patchGuestDriveByIdNoContent ", 204)
}

func (o *PatchGuestDriveByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGuestDriveByIDBadRequest creates a PatchGuestDriveByIDBadRequest with default headers values
func NewPatchGuestDriveByIDBadRequest() *PatchGuestDriveByIDBadRequest {
	return &PatchGuestDriveByIDBadRequest{}
}

/*PatchGuestDriveByIDBadRequest handles this case with default header values.

Drive cannot be updated due to bad input
*/
type PatchGuestDriveByIDBadRequest struct {
	Payload *models.Error
}

func (o *PatchGuestDriveByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /drives/{drive_id}][%d] patchGuestDriveByIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchGuestDriveByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGuestDriveByIDDefault creates a PatchGuestDriveByIDDefault with default headers values
func NewPatchGuestDriveByIDDefault(code int) *PatchGuestDriveByIDDefault {
	return &PatchGuestDriveByIDDefault{
		_statusCode: code,
	}
}

/*PatchGuestDriveByIDDefault handles this case with default header values.

Internal server error.
*/
type PatchGuestDriveByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch guest drive by ID default response
func (o *PatchGuestDriveByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchGuestDriveByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /drives/{drive_id}][%d] patchGuestDriveByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchGuestDriveByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
