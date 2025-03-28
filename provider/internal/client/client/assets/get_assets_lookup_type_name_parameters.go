// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetAssetsLookupTypeNameParams creates a new GetAssetsLookupTypeNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAssetsLookupTypeNameParams() *GetAssetsLookupTypeNameParams {
	return &GetAssetsLookupTypeNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAssetsLookupTypeNameParamsWithTimeout creates a new GetAssetsLookupTypeNameParams object
// with the ability to set a timeout on a request.
func NewGetAssetsLookupTypeNameParamsWithTimeout(timeout time.Duration) *GetAssetsLookupTypeNameParams {
	return &GetAssetsLookupTypeNameParams{
		timeout: timeout,
	}
}

// NewGetAssetsLookupTypeNameParamsWithContext creates a new GetAssetsLookupTypeNameParams object
// with the ability to set a context for a request.
func NewGetAssetsLookupTypeNameParamsWithContext(ctx context.Context) *GetAssetsLookupTypeNameParams {
	return &GetAssetsLookupTypeNameParams{
		Context: ctx,
	}
}

// NewGetAssetsLookupTypeNameParamsWithHTTPClient creates a new GetAssetsLookupTypeNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAssetsLookupTypeNameParamsWithHTTPClient(client *http.Client) *GetAssetsLookupTypeNameParams {
	return &GetAssetsLookupTypeNameParams{
		HTTPClient: client,
	}
}

/*
GetAssetsLookupTypeNameParams contains all the parameters to send to the API endpoint

	for the get assets lookup type name operation.

	Typically these are written to a http.Request.
*/
type GetAssetsLookupTypeNameParams struct {

	/* Name.

	   Asset name
	*/
	Name string

	/* Type.

	   Asset type
	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get assets lookup type name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAssetsLookupTypeNameParams) WithDefaults() *GetAssetsLookupTypeNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get assets lookup type name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAssetsLookupTypeNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get assets lookup type name params
func (o *GetAssetsLookupTypeNameParams) WithTimeout(timeout time.Duration) *GetAssetsLookupTypeNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get assets lookup type name params
func (o *GetAssetsLookupTypeNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get assets lookup type name params
func (o *GetAssetsLookupTypeNameParams) WithContext(ctx context.Context) *GetAssetsLookupTypeNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get assets lookup type name params
func (o *GetAssetsLookupTypeNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get assets lookup type name params
func (o *GetAssetsLookupTypeNameParams) WithHTTPClient(client *http.Client) *GetAssetsLookupTypeNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get assets lookup type name params
func (o *GetAssetsLookupTypeNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get assets lookup type name params
func (o *GetAssetsLookupTypeNameParams) WithName(name string) *GetAssetsLookupTypeNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get assets lookup type name params
func (o *GetAssetsLookupTypeNameParams) SetName(name string) {
	o.Name = name
}

// WithType adds the typeVar to the get assets lookup type name params
func (o *GetAssetsLookupTypeNameParams) WithType(typeVar string) *GetAssetsLookupTypeNameParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get assets lookup type name params
func (o *GetAssetsLookupTypeNameParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetAssetsLookupTypeNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
