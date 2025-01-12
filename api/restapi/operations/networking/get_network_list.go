// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetNetworkListHandlerFunc turns a function with the right signature into a get network list handler
type GetNetworkListHandlerFunc func(GetNetworkListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNetworkListHandlerFunc) Handle(params GetNetworkListParams) middleware.Responder {
	return fn(params)
}

// GetNetworkListHandler interface for that can handle valid get network list params
type GetNetworkListHandler interface {
	Handle(GetNetworkListParams) middleware.Responder
}

// NewGetNetworkList creates a new http.Handler for the get network list operation
func NewGetNetworkList(ctx *middleware.Context, handler GetNetworkListHandler) *GetNetworkList {
	return &GetNetworkList{Context: ctx, Handler: handler}
}

/*GetNetworkList swagger:route GET /networking networking getNetworkList

Get all networks

*/
type GetNetworkList struct {
	Context *middleware.Context
	Handler GetNetworkListHandler
}

func (o *GetNetworkList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNetworkListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
