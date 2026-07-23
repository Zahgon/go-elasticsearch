package cloneapikey

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	ApiKey string `json:"api_key"`

	Expiration types.Duration `json:"expiration,omitempty"`

	Metadata types.Metadata `json:"metadata,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
