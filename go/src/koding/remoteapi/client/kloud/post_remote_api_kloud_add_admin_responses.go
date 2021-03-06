package kloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIKloudAddAdminReader is a Reader for the PostRemoteAPIKloudAddAdmin structure.
type PostRemoteAPIKloudAddAdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIKloudAddAdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIKloudAddAdminOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIKloudAddAdminUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIKloudAddAdminOK creates a PostRemoteAPIKloudAddAdminOK with default headers values
func NewPostRemoteAPIKloudAddAdminOK() *PostRemoteAPIKloudAddAdminOK {
	return &PostRemoteAPIKloudAddAdminOK{}
}

/*PostRemoteAPIKloudAddAdminOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIKloudAddAdminOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIKloudAddAdminOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/Kloud.addAdmin][%d] postRemoteApiKloudAddAdminOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIKloudAddAdminOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIKloudAddAdminUnauthorized creates a PostRemoteAPIKloudAddAdminUnauthorized with default headers values
func NewPostRemoteAPIKloudAddAdminUnauthorized() *PostRemoteAPIKloudAddAdminUnauthorized {
	return &PostRemoteAPIKloudAddAdminUnauthorized{}
}

/*PostRemoteAPIKloudAddAdminUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIKloudAddAdminUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIKloudAddAdminUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/Kloud.addAdmin][%d] postRemoteApiKloudAddAdminUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIKloudAddAdminUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
