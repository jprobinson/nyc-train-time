package dialogflow

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type (
	// GoogleActionService returns a GoogleActionHandler for each action in their
	// Google Action's DialogFlow.
	GoogleActionService interface {
		// action name => handler
		Actions() map[string]GoogleActionHandler
	}

	// GoogleActionHandler encapsulates the logic for a single action for a Google
	// Action.
	// For more information about the request and response, see the DialogFlow
	// documentation for fulfillment: https://dialogflow.com/docs/fulfillment
	GoogleActionHandler func(context.Context, *GoogleRequest) (*GoogleFulfillmentResponse, error)
)

func (s service) postGoogle(ctx context.Context, req interface{}) (interface{}, error) {
	r, ok := req.(*GoogleRequest)
	if !ok {
		return nil, errBadRequest
	}
	handler, ok := s.google[r.Result.Action]
	if !ok {
		return nil, errBadRequest
	}
	return handler(ctx, r)
}

func decodeGoogle(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GoogleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errBadRequest
	}
	return &req, nil
}

// GoogleFulfillmentResponse is the response DialogFlow expects for Google Actions.
// More information here: https://dialogflow.com/docs/fulfillment#response
type GoogleFulfillmentResponse struct {
	Speech        string                   `json:"speech,omitempty"`
	DisplayText   string                   `json:"displayText,omitempty"`
	Source        string                   `json:"source,omitempty"`
	ContextOut    []map[string]interface{} `json:"contextOut"`
	FollowUpEvent *FollowUpEvent           `json:"followupEvent"`
}

// FollowUpEvent allows an action to redirect to a new dialogflow 'event' instead of
// immediately responding.
type FollowUpEvent struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

// GoogleRequest contains the information of an incoming DialogFlow request from Google
// Actions.
type GoogleRequest struct {
	OriginalRequest struct {
		Source  string `json:"source"`
		Version string `json:"version"`
		Data    struct {
			IsInSandbox bool `json:"isInSandbox"`
			Surface     struct {
				Capabilities []struct {
					Name string `json:"name"`
				} `json:"capabilities"`
			} `json:"surface"`
			Inputs []struct {
				RawInputs []struct {
					Query     string `json:"query"`
					InputType string `json:"inputType"`
				} `json:"rawInputs"`
				Arguments []struct {
					RawText   string `json:"rawText"`
					TextValue string `json:"textValue"`
					Name      string `json:"name"`
				} `json:"arguments"`
				Intent string `json:"intent"`
			} `json:"inputs"`
			User struct {
				LastSeen time.Time `json:"lastSeen"`
				Locale   string    `json:"locale"`
				UserID   string    `json:"userId"`
			} `json:"user"`
			Conversation struct {
				ConversationID    string `json:"conversationId"`
				Type              string `json:"type"`
				ConversationToken string `json:"conversationToken"`
			} `json:"conversation"`
			AvailableSurfaces []struct {
				Capabilities []struct {
					Name string `json:"name"`
				} `json:"capabilities"`
			} `json:"availableSurfaces"`
		} `json:"data"`
	} `json:"originalRequest"`
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Lang      string    `json:"lang"`
	Result    struct {
		Source           string                 `json:"source"`
		ResolvedQuery    string                 `json:"resolvedQuery"`
		Speech           string                 `json:"speech"`
		Action           string                 `json:"action"`
		ActionIncomplete bool                   `json:"actionIncomplete"`
		Parameters       map[string]interface{} `json:"parameters"`
		Contexts         []struct {
			Name       string                 `json:"name"`
			Parameters map[string]interface{} `json:"parameters"`
			Lifespan   int                    `json:"lifespan"`
		} `json:"contexts"`
		Metadata struct {
			MatchedParameters []struct {
				Required bool   `json:"required"`
				DataType string `json:"dataType"`
				Name     string `json:"name"`
				Value    string `json:"value"`
				Prompts  []struct {
					Lang  string `json:"lang"`
					Value string `json:"value"`
				} `json:"prompts"`
				IsList bool `json:"isList"`
			} `json:"matchedParameters"`
			IntentName                string `json:"intentName"`
			IntentID                  string `json:"intentId"`
			WebhookUsed               string `json:"webhookUsed"`
			WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
			NluResponseTime           int    `json:"nluResponseTime"`
		} `json:"metadata"`
		Fulfillment struct {
			Speech   string        `json:"speech"`
			Messages []interface{} `json:"messages"`
		} `json:"fulfillment"`
		Score float64 `json:"score"`
	} `json:"result"`
	Status struct {
		Code            int    `json:"code"`
		ErrorType       string `json:"errorType"`
		WebhookTimedOut bool   `json:"webhookTimedOut"`
	} `json:"status"`
	SessionID string `json:"sessionId"`
}
