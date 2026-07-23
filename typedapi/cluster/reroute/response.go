package reroute

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Acknowledged bool                       `json:"acknowledged"`
	Explanations []types.RerouteExplanation `json:"explanations,omitempty"`

	State json.RawMessage `json:"state,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
