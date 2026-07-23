package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/credentialmanagedby"
)

type AuthenticateApiKey struct {
	Id        string                                  `json:"id"`
	Internal  *bool                                   `json:"internal,omitempty"`
	ManagedBy credentialmanagedby.CredentialManagedBy `json:"managed_by"`
	Name      *string                                 `json:"name,omitempty"`
}

func (s *AuthenticateApiKey) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAuthenticateApiKey() *AuthenticateApiKey { _ = "STUB: not implemented"; return nil }
