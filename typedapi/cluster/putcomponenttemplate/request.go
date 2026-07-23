package putcomponenttemplate

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Deprecated *bool `json:"deprecated,omitempty"`

	Meta_ types.Metadata `json:"_meta,omitempty"`

	Template types.IndexTemplateMapping `json:"template"`

	Version *int64 `json:"version,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
