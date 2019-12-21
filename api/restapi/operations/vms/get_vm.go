// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetVMHandlerFunc turns a function with the right signature into a get VM handler
type GetVMHandlerFunc func(GetVMParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetVMHandlerFunc) Handle(params GetVMParams) middleware.Responder {
	return fn(params)
}

// GetVMHandler interface for that can handle valid get VM params
type GetVMHandler interface {
	Handle(GetVMParams) middleware.Responder
}

// NewGetVM creates a new http.Handler for the get VM operation
func NewGetVM(ctx *middleware.Context, handler GetVMHandler) *GetVM {
	return &GetVM{Context: ctx, Handler: handler}
}

/*GetVM swagger:route GET /vms/{vmID} vms getVm

Return a VM instance

Returns an isntance of VM

*/
type GetVM struct {
	Context *middleware.Context
	Handler GetVMHandler
}

func (o *GetVM) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetVMParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}