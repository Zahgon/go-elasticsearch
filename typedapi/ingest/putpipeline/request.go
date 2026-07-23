package putpipeline

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fieldaccesspattern"
)

type Request struct {
	Deprecated *bool `json:"deprecated,omitempty"`

	Description *string `json:"description,omitempty"`

	FieldAccessPattern *fieldaccesspattern.FieldAccessPattern `json:"field_access_pattern,omitempty"`

	Meta_ types.Metadata `json:"_meta,omitempty"`

	OnFailure []types.ProcessorContainer `json:"on_failure,omitempty"`

	Processors []types.ProcessorContainer `json:"processors,omitempty"`

	Version *int64 `json:"version,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
