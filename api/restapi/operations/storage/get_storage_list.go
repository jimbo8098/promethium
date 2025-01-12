// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetStorageListHandlerFunc turns a function with the right signature into a get storage list handler
type GetStorageListHandlerFunc func(GetStorageListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStorageListHandlerFunc) Handle(params GetStorageListParams) middleware.Responder {
	return fn(params)
}

// GetStorageListHandler interface for that can handle valid get storage list params
type GetStorageListHandler interface {
	Handle(GetStorageListParams) middleware.Responder
}

// NewGetStorageList creates a new http.Handler for the get storage list operation
func NewGetStorageList(ctx *middleware.Context, handler GetStorageListHandler) *GetStorageList {
	return &GetStorageList{Context: ctx, Handler: handler}
}

/*GetStorageList swagger:route GET /storage storage getStorageList

Get all storage

*/
type GetStorageList struct {
	Context *middleware.Context
	Handler GetStorageListHandler
}

func (o *GetStorageList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetStorageListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
