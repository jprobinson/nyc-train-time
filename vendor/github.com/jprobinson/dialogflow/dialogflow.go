package dialogflow

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/go-kit/kit/endpoint"
)

type (
	// FulfillmentService returns a IntentHandler for each intent in their DialogFlow.
	FulfillmentService interface {
		// Intents returns a mapping of intent display name IntentHandler.
		Intents() map[string]IntentHandler

		// Middleware can be used for things like adding generic responses or
		// authentication.
		Middleware(endpoint.Endpoint) endpoint.Endpoint
	}

	// IntentHandler encapsulates the logic for a single action for a Dialogflow intent.
	// For more information about the request and response, see the DialogFlow
	// documentation for fulfillment: https://dialogflow.com/docs/fulfillment/how-it-works
	IntentHandler func(context.Context, *Request) (*FulfillmentResponse, error)
)

func (s service) post(ctx context.Context, req interface{}) (interface{}, error) {
	r, ok := req.(*Request)
	if !ok {
		return nil, errBadRequest
	}
	handler, ok := s.intents[r.QueryResult.Intent.DisplayName]
	if !ok {
		return nil, errBadRequest
	}
	return handler(ctx, r)
}

func decode(ctx context.Context, r *http.Request) (interface{}, error) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		kit.LogErrorMsg(ctx, err, "unable to parse request")
		return nil, errBadRequest
	}
	defer r.Body.Close()
	return &req, nil
}

// GetParam is a convenience method for getting a parameter from QueryResult.Parameters.
// If the parameter does not exist, an error will be returned.
func (r *Request) GetParam(name string) (interface{}, error) {
	val, ok := r.QueryResult.Parameters[name]
	if !ok {
		return nil, errors.New("param not found")
	}
	return val, nil
}

// GetStringParam is a convenience method for getting a parameter from QueryResult.Parameters.
// If the parameter does not exist or is not a string, an error will be returned.
func (r *Request) GetStringParam(name string) (string, error) {
	val, err := r.GetParam(name)
	if err != nil {
		return "", err
	}

	str, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("param is of type %T, not string", val)
	}

	return str, nil
}

// Request contains all the inbound information from a Dialogflow fulfillment request.
// The QueryResult field contains most of the needed information. For more information
// about how to fill this struct, see the documentation at:
// https://dialogflow.com/docs/fulfillment/how-it-works.
type Request struct {
	ResponseID                  string      `json:"responseId"`
	Session                     string      `json:"session"`
	QueryResult                 QueryResult `json:"queryResult"`
	OriginalDetectIntentRequest struct {
		Source  string `json:"source"`
		Version string `json:"version"`
		Payload struct {
			IsInSandbox bool `json:"isInSandbox"`
			Surface     struct {
				Capabilities []interface{} `json:"capabilities"`
			} `json:"surface"`
			Inputs []interface{} `json:"inputs"`
			User   struct {
				UserID   string `json:"userId"`
				Locale   string `json:"locale"`
				LastSeen string `json:"lastSeen"`
			} `json:"user"`
			Conversation      interface{}   `json:"conversation"`
			AvailableSurfaces []interface{} `json:"availableSurfaces"`
		} `json:"payload"`
	} `json:"originalDetectIntentRequest"`
}

// QueryResult contains the result of the conversation query or event processing.
type QueryResult struct {
	LanguageCode             string                 `json:"languageCode"`
	QueryText                string                 `json:"queryText"`
	FulfillmentText          string                 `json:"fulfillmentText"`
	Action                   string                 `json:"action"`
	AllRequiredParamsPresent bool                   `json:"allRequiredParamsPresent"`
	Parameters               map[string]interface{} `json:"parameters"`
	OutputContexts           []OutputContext        `json:"outputContexts"`
	Intent                   Intent                 `json:"intent"`
	FulfillmentMessages      []struct {
		Text struct {
			Text []string `json:"text"`
		} `json:"text"`
	} `json:"fulfillmentMessages"`
	IntentDetectionConfidence float64     `json:"intentDetectionConfidence"`
	DiagnosticInfo            interface{} `json:"diagnosticInfo"`
}

// Intent contains the name that matched the user's query.
type Intent struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

// FulfillmentResponse contains all the fields required for responding to a Dialogflow
// intent. For more information about how to fill this struct, see the documentation at:
// https://dialogflow.com/docs/fulfillment/how-it-works.
type FulfillmentResponse struct {
	FulfillmentText     string               `json:"fulfillmentText"`
	FulfillmentMessages []FulfillmentMessage `json:"fulfillmentMessages"`
	Source              string               `json:"source"`
	Payload             *FulfillmentPayload  `json:"payload"`
	OutputContexts      []OutputContext      `json:"outputContexts"`
	FollowupEventInput  *EventInput          `json:"followupEventInput"`
}

// FulfillmentMessage is the most basic way to respond to intents.
type FulfillmentMessage struct {
	Text []string `json:"text"`
}

// FulfillmentPayload allows users to respond differently to alternate platforms.
type FulfillmentPayload struct {
	Google   RichFulfillmentPayload `json:"google"`
	Facebook FulfillmentMessage     `json:"facebook"`
	Slack    FulfillmentMessage     `json:"slack"`
}

// RichFulfillmentPayload can be used for rich responses to Google Actions.
type RichFulfillmentPayload struct {
	ExpectUserResponse bool `json:"expectUserResponse"`
	RichResponse       struct {
		Items []interface{} `json:"items"`
	} `json:"richResponse"`
}

// OutputContext holds information for output contexts within fulfillment requests and
// responses.
type OutputContext struct {
	Name          string                 `json:"name"`
	LifespanCount int                    `json:"lifespanCount"`
	Parameters    map[string]interface{} `json:"parameters"`
}

// EventInput will hold all inbound input information.
type EventInput struct {
	Name         string                 `json:"name"`
	LanguageCode string                 `json:"languageCode"`
	Parameters   map[string]interface{} `json:"parameters"`
}
