package termvectors

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type Request struct {
	Doc json.RawMessage `json:"doc,omitempty"`

	FieldStatistics *bool `json:"field_statistics,omitempty"`

	Fields []string `json:"fields,omitempty"`

	Filter *types.TermVectorsFilter `json:"filter,omitempty"`

	Offsets *bool `json:"offsets,omitempty"`

	Payloads *bool `json:"payloads,omitempty"`

	PerFieldAnalyzer map[string]string `json:"per_field_analyzer,omitempty"`

	Positions *bool `json:"positions,omitempty"`

	Routing []string `json:"routing,omitempty"`

	TermStatistics *bool `json:"term_statistics,omitempty"`

	Version *int64 `json:"version,omitempty"`

	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
