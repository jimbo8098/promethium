// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteVMHandlerFunc turns a function with the right signature into a delete VM handler
type DeleteVMHandlerFunc func(DeleteVMParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteVMHandlerFunc) Handle(params DeleteVMParams) middleware.Responder {
	return fn(params)
}

// DeleteVMHandler interface for that can handle valid delete VM params
type DeleteVMHandler interface {
	Handle(DeleteVMParams) middleware.Responder
}

// NewDeleteVM creates a new http.Handler for the delete VM operation
func NewDeleteVM(ctx *middleware.Context, handler DeleteVMHandler) *DeleteVM {
	return &DeleteVM{Context: ctx, Handler: handler}
}

/*DeleteVM swagger:route DELETE /vms/{vmID} vms deleteVm

Destroy a VM instance

Destroy an isntance of VM

*/
type DeleteVM struct {
	Context *middleware.Context
	Handler DeleteVMHandler
}

func (o *DeleteVM) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteVMParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
