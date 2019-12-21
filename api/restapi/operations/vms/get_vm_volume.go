// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetVMVolumeHandlerFunc turns a function with the right signature into a get VM volume handler
type GetVMVolumeHandlerFunc func(GetVMVolumeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetVMVolumeHandlerFunc) Handle(params GetVMVolumeParams) middleware.Responder {
	return fn(params)
}

// GetVMVolumeHandler interface for that can handle valid get VM volume params
type GetVMVolumeHandler interface {
	Handle(GetVMVolumeParams) middleware.Responder
}

// NewGetVMVolume creates a new http.Handler for the get VM volume operation
func NewGetVMVolume(ctx *middleware.Context, handler GetVMVolumeHandler) *GetVMVolume {
	return &GetVMVolume{Context: ctx, Handler: handler}
}

/*GetVMVolume swagger:route GET /vms/{vmID}/volumes/{volumeID} vms getVmVolume

Return a VM instance

Returns an isntance of VM

*/
type GetVMVolume struct {
	Context *middleware.Context
	Handler GetVMVolumeHandler
}

func (o *GetVMVolume) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetVMVolumeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
