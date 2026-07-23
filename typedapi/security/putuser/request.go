package putuser

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Email *string `json:"email,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	FullName *string `json:"full_name,omitempty"`

	Metadata types.Metadata `json:"metadata,omitempty"`

	Password *string `json:"password,omitempty"`

	PasswordHash *string `json:"password_hash,omitempty"`

	Roles    []string `json:"roles,omitempty"`
	Username *string  `json:"username,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
