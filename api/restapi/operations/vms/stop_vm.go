// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// StopVMHandlerFunc turns a function with the right signature into a stop VM handler
type StopVMHandlerFunc func(StopVMParams) middleware.Responder

// Handle executing the request and returning a response
func (fn StopVMHandlerFunc) Handle(params StopVMParams) middleware.Responder {
	return fn(params)
}

// StopVMHandler interface for that can handle valid stop VM params
type StopVMHandler interface {
	Handle(StopVMParams) middleware.Responder
}

// NewStopVM creates a new http.Handler for the stop VM operation
func NewStopVM(ctx *middleware.Context, handler StopVMHandler) *StopVM {
	return &StopVM{Context: ctx, Handler: handler}
}

/*StopVM swagger:route GET /vms/{vmID}/stop vms stopVm

Stop a VM instance

Stops an isntance of VM

*/
type StopVM struct {
	Context *middleware.Context
	Handler StopVMHandler
}

func (o *StopVM) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewStopVMParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
