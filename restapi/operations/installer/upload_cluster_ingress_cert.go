// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UploadClusterIngressCertHandlerFunc turns a function with the right signature into a upload cluster ingress cert handler
type UploadClusterIngressCertHandlerFunc func(UploadClusterIngressCertParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UploadClusterIngressCertHandlerFunc) Handle(params UploadClusterIngressCertParams) middleware.Responder {
	return fn(params)
}

// UploadClusterIngressCertHandler interface for that can handle valid upload cluster ingress cert params
type UploadClusterIngressCertHandler interface {
	Handle(UploadClusterIngressCertParams) middleware.Responder
}

// NewUploadClusterIngressCert creates a new http.Handler for the upload cluster ingress cert operation
func NewUploadClusterIngressCert(ctx *middleware.Context, handler UploadClusterIngressCertHandler) *UploadClusterIngressCert {
	return &UploadClusterIngressCert{Context: ctx, Handler: handler}
}

/*UploadClusterIngressCert swagger:route POST /clusters/{cluster_id}/uploads/ingress-cert installer uploadClusterIngressCert

Transfer the ingress certificate for the cluster.

*/
type UploadClusterIngressCert struct {
	Context *middleware.Context
	Handler UploadClusterIngressCertHandler
}

func (o *UploadClusterIngressCert) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUploadClusterIngressCertParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
