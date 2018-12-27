// Package dialogflow is a framework for building services that respond to DialogFlow
// fulfillment webhooks.
package dialogflow

import (
	"context"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"google.golang.org/grpc"
)

type service struct {
	intents    map[string]IntentHandler
	middleware endpoint.Middleware
}

// Run will register a service with kit and run the server. Call this in your main
// function.
//
// The service will register the webhook on the path "/fulfillment" so make sure to
// configure your fulfillment webhook to point at something like:
// https://example.appspot.com/fulfillment
func Run(svc FulfillmentService) error {
	return kit.Run(&service{intents: svc.Intents(), middleware: svc.Middleware})
}

func (s *service) HTTPOptions() []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			httptransport.EncodeJSONResponse(ctx, w, err)
		}),
	}
}

func (s *service) HTTPRouterOptions() []kit.RouterOption {
	return []kit.RouterOption{kit.RouterSelect("stdlib")}
}

func (s *service) HTTPMiddleware(h http.Handler) http.Handler {
	return h
}

func (s *service) Middleware(ep endpoint.Endpoint) endpoint.Endpoint {
	return s.middleware(ep)
}

func (s *service) HTTPEndpoints() map[string]map[string]kit.HTTPEndpoint {
	return map[string]map[string]kit.HTTPEndpoint{
		"/fulfillment": {
			"POST": {
				Endpoint: s.post,
				Decoder:  decode,
			},
		},
	}
}

var errBadRequest = kit.NewJSONStatusResponse(map[string]string{
	"error": "bad request"}, http.StatusBadRequest)

func (s *service) RPCMiddleware() grpc.UnaryServerInterceptor {
	return nil
}

func (s *service) RPCServiceDesc() *grpc.ServiceDesc {
	return nil
}

func (s *service) RPCOptions() []grpc.ServerOption {
	return nil
}
