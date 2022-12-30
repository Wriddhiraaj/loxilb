// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// DeleteConfigVlanVlanIDNoContentCode is the HTTP code returned for type DeleteConfigVlanVlanIDNoContent
const DeleteConfigVlanVlanIDNoContentCode int = 204

/*
DeleteConfigVlanVlanIDNoContent OK

swagger:response deleteConfigVlanVlanIdNoContent
*/
type DeleteConfigVlanVlanIDNoContent struct {
}

// NewDeleteConfigVlanVlanIDNoContent creates DeleteConfigVlanVlanIDNoContent with default headers values
func NewDeleteConfigVlanVlanIDNoContent() *DeleteConfigVlanVlanIDNoContent {

	return &DeleteConfigVlanVlanIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteConfigVlanVlanIDBadRequestCode is the HTTP code returned for type DeleteConfigVlanVlanIDBadRequest
const DeleteConfigVlanVlanIDBadRequestCode int = 400

/*
DeleteConfigVlanVlanIDBadRequest Malformed arguments for API call

swagger:response deleteConfigVlanVlanIdBadRequest
*/
type DeleteConfigVlanVlanIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigVlanVlanIDBadRequest creates DeleteConfigVlanVlanIDBadRequest with default headers values
func NewDeleteConfigVlanVlanIDBadRequest() *DeleteConfigVlanVlanIDBadRequest {

	return &DeleteConfigVlanVlanIDBadRequest{}
}

// WithPayload adds the payload to the delete config vlan vlan Id bad request response
func (o *DeleteConfigVlanVlanIDBadRequest) WithPayload(payload *models.Error) *DeleteConfigVlanVlanIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config vlan vlan Id bad request response
func (o *DeleteConfigVlanVlanIDBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigVlanVlanIDUnauthorizedCode is the HTTP code returned for type DeleteConfigVlanVlanIDUnauthorized
const DeleteConfigVlanVlanIDUnauthorizedCode int = 401

/*
DeleteConfigVlanVlanIDUnauthorized Invalid authentication credentials

swagger:response deleteConfigVlanVlanIdUnauthorized
*/
type DeleteConfigVlanVlanIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigVlanVlanIDUnauthorized creates DeleteConfigVlanVlanIDUnauthorized with default headers values
func NewDeleteConfigVlanVlanIDUnauthorized() *DeleteConfigVlanVlanIDUnauthorized {

	return &DeleteConfigVlanVlanIDUnauthorized{}
}

// WithPayload adds the payload to the delete config vlan vlan Id unauthorized response
func (o *DeleteConfigVlanVlanIDUnauthorized) WithPayload(payload *models.Error) *DeleteConfigVlanVlanIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config vlan vlan Id unauthorized response
func (o *DeleteConfigVlanVlanIDUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigVlanVlanIDForbiddenCode is the HTTP code returned for type DeleteConfigVlanVlanIDForbidden
const DeleteConfigVlanVlanIDForbiddenCode int = 403

/*
DeleteConfigVlanVlanIDForbidden Capacity insufficient

swagger:response deleteConfigVlanVlanIdForbidden
*/
type DeleteConfigVlanVlanIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigVlanVlanIDForbidden creates DeleteConfigVlanVlanIDForbidden with default headers values
func NewDeleteConfigVlanVlanIDForbidden() *DeleteConfigVlanVlanIDForbidden {

	return &DeleteConfigVlanVlanIDForbidden{}
}

// WithPayload adds the payload to the delete config vlan vlan Id forbidden response
func (o *DeleteConfigVlanVlanIDForbidden) WithPayload(payload *models.Error) *DeleteConfigVlanVlanIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config vlan vlan Id forbidden response
func (o *DeleteConfigVlanVlanIDForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigVlanVlanIDNotFoundCode is the HTTP code returned for type DeleteConfigVlanVlanIDNotFound
const DeleteConfigVlanVlanIDNotFoundCode int = 404

/*
DeleteConfigVlanVlanIDNotFound Resource not found

swagger:response deleteConfigVlanVlanIdNotFound
*/
type DeleteConfigVlanVlanIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigVlanVlanIDNotFound creates DeleteConfigVlanVlanIDNotFound with default headers values
func NewDeleteConfigVlanVlanIDNotFound() *DeleteConfigVlanVlanIDNotFound {

	return &DeleteConfigVlanVlanIDNotFound{}
}

// WithPayload adds the payload to the delete config vlan vlan Id not found response
func (o *DeleteConfigVlanVlanIDNotFound) WithPayload(payload *models.Error) *DeleteConfigVlanVlanIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config vlan vlan Id not found response
func (o *DeleteConfigVlanVlanIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigVlanVlanIDConflictCode is the HTTP code returned for type DeleteConfigVlanVlanIDConflict
const DeleteConfigVlanVlanIDConflictCode int = 409

/*
DeleteConfigVlanVlanIDConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response deleteConfigVlanVlanIdConflict
*/
type DeleteConfigVlanVlanIDConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigVlanVlanIDConflict creates DeleteConfigVlanVlanIDConflict with default headers values
func NewDeleteConfigVlanVlanIDConflict() *DeleteConfigVlanVlanIDConflict {

	return &DeleteConfigVlanVlanIDConflict{}
}

// WithPayload adds the payload to the delete config vlan vlan Id conflict response
func (o *DeleteConfigVlanVlanIDConflict) WithPayload(payload *models.Error) *DeleteConfigVlanVlanIDConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config vlan vlan Id conflict response
func (o *DeleteConfigVlanVlanIDConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigVlanVlanIDInternalServerErrorCode is the HTTP code returned for type DeleteConfigVlanVlanIDInternalServerError
const DeleteConfigVlanVlanIDInternalServerErrorCode int = 500

/*
DeleteConfigVlanVlanIDInternalServerError Internal service error

swagger:response deleteConfigVlanVlanIdInternalServerError
*/
type DeleteConfigVlanVlanIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigVlanVlanIDInternalServerError creates DeleteConfigVlanVlanIDInternalServerError with default headers values
func NewDeleteConfigVlanVlanIDInternalServerError() *DeleteConfigVlanVlanIDInternalServerError {

	return &DeleteConfigVlanVlanIDInternalServerError{}
}

// WithPayload adds the payload to the delete config vlan vlan Id internal server error response
func (o *DeleteConfigVlanVlanIDInternalServerError) WithPayload(payload *models.Error) *DeleteConfigVlanVlanIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config vlan vlan Id internal server error response
func (o *DeleteConfigVlanVlanIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigVlanVlanIDServiceUnavailableCode is the HTTP code returned for type DeleteConfigVlanVlanIDServiceUnavailable
const DeleteConfigVlanVlanIDServiceUnavailableCode int = 503

/*
DeleteConfigVlanVlanIDServiceUnavailable Maintanence mode

swagger:response deleteConfigVlanVlanIdServiceUnavailable
*/
type DeleteConfigVlanVlanIDServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigVlanVlanIDServiceUnavailable creates DeleteConfigVlanVlanIDServiceUnavailable with default headers values
func NewDeleteConfigVlanVlanIDServiceUnavailable() *DeleteConfigVlanVlanIDServiceUnavailable {

	return &DeleteConfigVlanVlanIDServiceUnavailable{}
}

// WithPayload adds the payload to the delete config vlan vlan Id service unavailable response
func (o *DeleteConfigVlanVlanIDServiceUnavailable) WithPayload(payload *models.Error) *DeleteConfigVlanVlanIDServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config vlan vlan Id service unavailable response
func (o *DeleteConfigVlanVlanIDServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}