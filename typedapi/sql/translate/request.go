package translate

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	FetchSize *int `json:"fetch_size,omitempty"`

	Filter *types.Query `json:"filter,omitempty"`

	Query string `json:"query"`

	TimeZone *string `json:"time_zone,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
