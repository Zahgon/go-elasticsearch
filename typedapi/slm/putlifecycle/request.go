package putlifecycle

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Config *types.Configuration `json:"config,omitempty"`

	Name *string `json:"name,omitempty"`

	Repository *string `json:"repository,omitempty"`

	Retention *types.Retention `json:"retention,omitempty"`

	Schedule *string `json:"schedule,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
