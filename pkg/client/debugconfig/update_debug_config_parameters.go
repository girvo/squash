package debugconfig

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/solo-io/squash/pkg/models"
)

// NewUpdateDebugConfigParams creates a new UpdateDebugConfigParams object
// with the default values initialized.
func NewUpdateDebugConfigParams() *UpdateDebugConfigParams {
	var ()
	return &UpdateDebugConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDebugConfigParamsWithTimeout creates a new UpdateDebugConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDebugConfigParamsWithTimeout(timeout time.Duration) *UpdateDebugConfigParams {
	var ()
	return &UpdateDebugConfigParams{

		timeout: timeout,
	}
}

// NewUpdateDebugConfigParamsWithContext creates a new UpdateDebugConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDebugConfigParamsWithContext(ctx context.Context) *UpdateDebugConfigParams {
	var ()
	return &UpdateDebugConfigParams{

		Context: ctx,
	}
}

// NewUpdateDebugConfigParamsWithHTTPClient creates a new UpdateDebugConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDebugConfigParamsWithHTTPClient(client *http.Client) *UpdateDebugConfigParams {
	var ()
	return &UpdateDebugConfigParams{
		HTTPClient: client,
	}
}

/*UpdateDebugConfigParams contains all the parameters to send to the API endpoint
for the update debug config operation typically these are written to a http.Request
*/
type UpdateDebugConfigParams struct {

	/*Body
	  debug config edited

	*/
	Body *models.DebugConfig
	/*DebugConfigID
	  ID of config to return

	*/
	DebugConfigID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update debug config params
func (o *UpdateDebugConfigParams) WithTimeout(timeout time.Duration) *UpdateDebugConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update debug config params
func (o *UpdateDebugConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update debug config params
func (o *UpdateDebugConfigParams) WithContext(ctx context.Context) *UpdateDebugConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update debug config params
func (o *UpdateDebugConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update debug config params
func (o *UpdateDebugConfigParams) WithHTTPClient(client *http.Client) *UpdateDebugConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update debug config params
func (o *UpdateDebugConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update debug config params
func (o *UpdateDebugConfigParams) WithBody(body *models.DebugConfig) *UpdateDebugConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update debug config params
func (o *UpdateDebugConfigParams) SetBody(body *models.DebugConfig) {
	o.Body = body
}

// WithDebugConfigID adds the debugConfigID to the update debug config params
func (o *UpdateDebugConfigParams) WithDebugConfigID(debugConfigID string) *UpdateDebugConfigParams {
	o.SetDebugConfigID(debugConfigID)
	return o
}

// SetDebugConfigID adds the debugConfigId to the update debug config params
func (o *UpdateDebugConfigParams) SetDebugConfigID(debugConfigID string) {
	o.DebugConfigID = debugConfigID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDebugConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.DebugConfig)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param debugConfigId
	if err := r.SetPathParam("debugConfigId", o.DebugConfigID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
