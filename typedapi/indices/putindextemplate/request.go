package putindextemplate

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	AllowAutoCreate *bool `json:"allow_auto_create,omitempty"`

	ComposedOf []string `json:"composed_of,omitempty"`

	DataStream *types.DataStreamVisibility `json:"data_stream,omitempty"`

	Deprecated *bool `json:"deprecated,omitempty"`

	IgnoreMissingComponentTemplates []string `json:"ignore_missing_component_templates,omitempty"`

	IndexPatterns []string `json:"index_patterns,omitempty"`

	Meta_ types.Metadata `json:"_meta,omitempty"`

	Priority *int64 `json:"priority,omitempty"`

	Template *types.IndexTemplateMapping `json:"template,omitempty"`

	Version *int64 `json:"version,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
