package authenticate

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	ApiKey              *types.AuthenticateApiKey `json:"api_key,omitempty"`
	AuthenticationRealm types.RealmInfo           `json:"authentication_realm"`
	AuthenticationType  string                    `json:"authentication_type"`
	Email               *string                   `json:"email,omitempty"`
	Enabled             bool                      `json:"enabled"`
	FullName            *string                   `json:"full_name,omitempty"`
	LookupRealm         types.RealmInfo           `json:"lookup_realm"`
	Metadata            types.Metadata            `json:"metadata"`
	Roles               []string                  `json:"roles"`
	Token               *types.AuthenticateToken  `json:"token,omitempty"`
	Username            string                    `json:"username"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
