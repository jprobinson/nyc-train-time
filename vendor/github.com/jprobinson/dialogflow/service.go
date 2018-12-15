// A Google App Engine framework for building services that respond to DialogFlow actions.
package dialogflow

import (
	"context"
	"net/http"

	"github.com/NYTimes/marvin"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type service struct {
	google     map[string]GoogleActionHandler
	middleware endpoint.Middleware
}

// Init will register a service with marvin and Google App Engine.
// Call this in your init function or main function just before appengine.Main.
//
// The service will register the webhook on the path "/google" so make sure to configure
// your fulfillment webhook to point at something like https://example.com/google
func Init(google GoogleActionService, middleware endpoint.Middleware) {
	marvin.Init(&service{google: google.Actions(), middleware: middleware})
}

func (s *service) Options() []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			httptransport.EncodeJSONResponse(ctx, w, err)
		}),
	}
}

func (s *service) RouterOptions() []marvin.RouterOption {
	return []marvin.RouterOption{marvin.RouterSelect("stdlib")}
}

func (s *service) HTTPMiddleware(h http.Handler) http.Handler {
	return h
}

func (s *service) Middleware(ep endpoint.Endpoint) endpoint.Endpoint {
	if s.middleware != nil {
		return s.middleware(ep)
	}
	return ep
}

func (s *service) JSONEndpoints() map[string]map[string]marvin.HTTPEndpoint {
	return map[string]map[string]marvin.HTTPEndpoint{
		"/google": {
			"POST": {
				Endpoint: s.postGoogle,
				Decoder:  decodeGoogle,
			},
		},
	}
}

var errBadRequest = marvin.NewJSONStatusResponse(map[string]string{
	"error": "bad request"}, http.StatusBadRequest)
