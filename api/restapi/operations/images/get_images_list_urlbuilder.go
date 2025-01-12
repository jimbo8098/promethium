// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetImagesListURL generates an URL for the get images list operation
type GetImagesListURL struct {
	Limit *int32
	Since *int64
	Skip  *int32

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetImagesListURL) WithBasePath(bp string) *GetImagesListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetImagesListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetImagesListURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/images"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt32(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var sinceQ string
	if o.Since != nil {
		sinceQ = swag.FormatInt64(*o.Since)
	}
	if sinceQ != "" {
		qs.Set("since", sinceQ)
	}

	var skipQ string
	if o.Skip != nil {
		skipQ = swag.FormatInt32(*o.Skip)
	}
	if skipQ != "" {
		qs.Set("skip", skipQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetImagesListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetImagesListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetImagesListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetImagesListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetImagesListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetImagesListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
