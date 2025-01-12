// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateVMVolumeHandlerFunc turns a function with the right signature into a update VM volume handler
type UpdateVMVolumeHandlerFunc func(UpdateVMVolumeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateVMVolumeHandlerFunc) Handle(params UpdateVMVolumeParams) middleware.Responder {
	return fn(params)
}

// UpdateVMVolumeHandler interface for that can handle valid update VM volume params
type UpdateVMVolumeHandler interface {
	Handle(UpdateVMVolumeParams) middleware.Responder
}

// NewUpdateVMVolume creates a new http.Handler for the update VM volume operation
func NewUpdateVMVolume(ctx *middleware.Context, handler UpdateVMVolumeHandler) *UpdateVMVolume {
	return &UpdateVMVolume{Context: ctx, Handler: handler}
}

/*UpdateVMVolume swagger:route PUT /vms/{vmID}/volumes/{volumeID} vms updateVmVolume

Update a VM Interface instance

Update an instance of VM interface

*/
type UpdateVMVolume struct {
	Context *middleware.Context
	Handler UpdateVMVolumeHandler
}

func (o *UpdateVMVolume) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateVMVolumeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
