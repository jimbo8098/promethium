// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/768bit/promethium/api/models"
)

// NewUpdateStorageParams creates a new UpdateStorageParams object
// with the default values initialized.
func NewUpdateStorageParams() *UpdateStorageParams {
	var ()
	return &UpdateStorageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateStorageParamsWithTimeout creates a new UpdateStorageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateStorageParamsWithTimeout(timeout time.Duration) *UpdateStorageParams {
	var ()
	return &UpdateStorageParams{

		timeout: timeout,
	}
}

// NewUpdateStorageParamsWithContext creates a new UpdateStorageParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateStorageParamsWithContext(ctx context.Context) *UpdateStorageParams {
	var ()
	return &UpdateStorageParams{

		Context: ctx,
	}
}

// NewUpdateStorageParamsWithHTTPClient creates a new UpdateStorageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateStorageParamsWithHTTPClient(client *http.Client) *UpdateStorageParams {
	var ()
	return &UpdateStorageParams{
		HTTPClient: client,
	}
}

/*UpdateStorageParams contains all the parameters to send to the API endpoint
for the update storage operation typically these are written to a http.Request
*/
type UpdateStorageParams struct {

	/*StorageConfig
	  Create new VM instance

	*/
	StorageConfig models.UpdateStorage
	/*StorageID
	  ID of VM to return

	*/
	StorageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update storage params
func (o *UpdateStorageParams) WithTimeout(timeout time.Duration) *UpdateStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update storage params
func (o *UpdateStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update storage params
func (o *UpdateStorageParams) WithContext(ctx context.Context) *UpdateStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update storage params
func (o *UpdateStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update storage params
func (o *UpdateStorageParams) WithHTTPClient(client *http.Client) *UpdateStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update storage params
func (o *UpdateStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStorageConfig adds the storageConfig to the update storage params
func (o *UpdateStorageParams) WithStorageConfig(storageConfig models.UpdateStorage) *UpdateStorageParams {
	o.SetStorageConfig(storageConfig)
	return o
}

// SetStorageConfig adds the storageConfig to the update storage params
func (o *UpdateStorageParams) SetStorageConfig(storageConfig models.UpdateStorage) {
	o.StorageConfig = storageConfig
}

// WithStorageID adds the storageID to the update storage params
func (o *UpdateStorageParams) WithStorageID(storageID string) *UpdateStorageParams {
	o.SetStorageID(storageID)
	return o
}

// SetStorageID adds the storageId to the update storage params
func (o *UpdateStorageParams) SetStorageID(storageID string) {
	o.StorageID = storageID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.StorageConfig != nil {
		if err := r.SetBodyParam(o.StorageConfig); err != nil {
			return err
		}
	}

	// path param storageID
	if err := r.SetPathParam("storageID", o.StorageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
