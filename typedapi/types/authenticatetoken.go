package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/credentialmanagedby"
)

type AuthenticateToken struct {
	ManagedBy *credentialmanagedby.CredentialManagedBy `json:"managed_by,omitempty"`
	Name      *string                                  `json:"name,omitempty"`
	Type      *string                                  `json:"type,omitempty"`
}

func (s *AuthenticateToken) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAuthenticateToken() *AuthenticateToken { _ = "STUB: not implemented"; return nil }
