// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateVMHandlerFunc turns a function with the right signature into a update VM handler
type UpdateVMHandlerFunc func(UpdateVMParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateVMHandlerFunc) Handle(params UpdateVMParams) middleware.Responder {
	return fn(params)
}

// UpdateVMHandler interface for that can handle valid update VM params
type UpdateVMHandler interface {
	Handle(UpdateVMParams) middleware.Responder
}

// NewUpdateVM creates a new http.Handler for the update VM operation
func NewUpdateVM(ctx *middleware.Context, handler UpdateVMHandler) *UpdateVM {
	return &UpdateVM{Context: ctx, Handler: handler}
}

/*UpdateVM swagger:route PUT /vms/{vmID} vms updateVm

Update a VM instance

Update an instance of VM

*/
type UpdateVM struct {
	Context *middleware.Context
	Handler UpdateVMHandler
}

func (o *UpdateVM) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateVMParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
