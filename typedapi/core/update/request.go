package update

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	DetectNoop *bool `json:"detect_noop,omitempty"`

	Doc json.RawMessage `json:"doc,omitempty"`

	DocAsUpsert *bool `json:"doc_as_upsert,omitempty"`

	Script *types.Script `json:"script,omitempty"`

	ScriptedUpsert *bool `json:"scripted_upsert,omitempty"`

	Source_ types.SourceConfig `json:"_source,omitempty"`

	Upsert json.RawMessage `json:"upsert,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
