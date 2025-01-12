// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateVMHandlerFunc turns a function with the right signature into a create VM handler
type CreateVMHandlerFunc func(CreateVMParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateVMHandlerFunc) Handle(params CreateVMParams) middleware.Responder {
	return fn(params)
}

// CreateVMHandler interface for that can handle valid create VM params
type CreateVMHandler interface {
	Handle(CreateVMParams) middleware.Responder
}

// NewCreateVM creates a new http.Handler for the create VM operation
func NewCreateVM(ctx *middleware.Context, handler CreateVMHandler) *CreateVM {
	return &CreateVM{Context: ctx, Handler: handler}
}

/*CreateVM swagger:route POST /vms vms createVm

Create a VM instance

Create an instance of VM

*/
type CreateVM struct {
	Context *middleware.Context
	Handler CreateVMHandler
}

func (o *CreateVM) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateVMParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
