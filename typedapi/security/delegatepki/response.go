package delegatepki

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	AccessToken    string                `json:"access_token"`
	Authentication *types.Authentication `json:"authentication,omitempty"`

	ExpiresIn int64 `json:"expires_in"`

	Type string `json:"type"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
