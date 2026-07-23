package gettoken

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	AccessToken                         string                  `json:"access_token"`
	Authentication                      types.AuthenticatedUser `json:"authentication"`
	ExpiresIn                           int64                   `json:"expires_in"`
	KerberosAuthenticationResponseToken *string                 `json:"kerberos_authentication_response_token,omitempty"`
	RefreshToken                        *string                 `json:"refresh_token,omitempty"`
	Scope                               *string                 `json:"scope,omitempty"`
	Type                                string                  `json:"type"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
