// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigFdbAllHandlerFunc turns a function with the right signature into a get config fdb all handler
type GetConfigFdbAllHandlerFunc func(GetConfigFdbAllParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigFdbAllHandlerFunc) Handle(params GetConfigFdbAllParams) middleware.Responder {
	return fn(params)
}

// GetConfigFdbAllHandler interface for that can handle valid get config fdb all params
type GetConfigFdbAllHandler interface {
	Handle(GetConfigFdbAllParams) middleware.Responder
}

// NewGetConfigFdbAll creates a new http.Handler for the get config fdb all operation
func NewGetConfigFdbAll(ctx *middleware.Context, handler GetConfigFdbAllHandler) *GetConfigFdbAll {
	return &GetConfigFdbAll{Context: ctx, Handler: handler}
}

/*
	GetConfigFdbAll swagger:route GET /config/fdb/all getConfigFdbAll

Get FDB in the device(interface)

Get FDB in the device(interface).
*/
type GetConfigFdbAll struct {
	Context *middleware.Context
	Handler GetConfigFdbAllHandler
}

func (o *GetConfigFdbAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetConfigFdbAllParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetConfigFdbAllOKBody get config fdb all o k body
//
// swagger:model GetConfigFdbAllOKBody
type GetConfigFdbAllOKBody struct {

	// fdb attr
	FdbAttr []*models.FDBEntry `json:"fdbAttr"`
}

// Validate validates this get config fdb all o k body
func (o *GetConfigFdbAllOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFdbAttr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigFdbAllOKBody) validateFdbAttr(formats strfmt.Registry) error {
	if swag.IsZero(o.FdbAttr) { // not required
		return nil
	}

	for i := 0; i < len(o.FdbAttr); i++ {
		if swag.IsZero(o.FdbAttr[i]) { // not required
			continue
		}

		if o.FdbAttr[i] != nil {
			if err := o.FdbAttr[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigFdbAllOK" + "." + "fdbAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigFdbAllOK" + "." + "fdbAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get config fdb all o k body based on the context it is used
func (o *GetConfigFdbAllOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFdbAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigFdbAllOKBody) contextValidateFdbAttr(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.FdbAttr); i++ {

		if o.FdbAttr[i] != nil {
			if err := o.FdbAttr[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigFdbAllOK" + "." + "fdbAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigFdbAllOK" + "." + "fdbAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigFdbAllOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigFdbAllOKBody) UnmarshalBinary(b []byte) error {
	var res GetConfigFdbAllOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}