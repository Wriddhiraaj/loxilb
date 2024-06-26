// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostConfigBgpPolicyDefinitionsHandlerFunc turns a function with the right signature into a post config bgp policy definitions handler
type PostConfigBgpPolicyDefinitionsHandlerFunc func(PostConfigBgpPolicyDefinitionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostConfigBgpPolicyDefinitionsHandlerFunc) Handle(params PostConfigBgpPolicyDefinitionsParams) middleware.Responder {
	return fn(params)
}

// PostConfigBgpPolicyDefinitionsHandler interface for that can handle valid post config bgp policy definitions params
type PostConfigBgpPolicyDefinitionsHandler interface {
	Handle(PostConfigBgpPolicyDefinitionsParams) middleware.Responder
}

// NewPostConfigBgpPolicyDefinitions creates a new http.Handler for the post config bgp policy definitions operation
func NewPostConfigBgpPolicyDefinitions(ctx *middleware.Context, handler PostConfigBgpPolicyDefinitionsHandler) *PostConfigBgpPolicyDefinitions {
	return &PostConfigBgpPolicyDefinitions{Context: ctx, Handler: handler}
}

/*
	PostConfigBgpPolicyDefinitions swagger:route POST /config/bgp/policy/definitions postConfigBgpPolicyDefinitions

# Adds a BGP Policy

Adds a BGP Policy
*/
type PostConfigBgpPolicyDefinitions struct {
	Context *middleware.Context
	Handler PostConfigBgpPolicyDefinitionsHandler
}

func (o *PostConfigBgpPolicyDefinitions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostConfigBgpPolicyDefinitionsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
