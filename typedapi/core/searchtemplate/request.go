package searchtemplate

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Explain *bool `json:"explain,omitempty"`

	Id *string `json:"id,omitempty"`

	Params map[string]json.RawMessage `json:"params,omitempty"`

	Profile *bool `json:"profile,omitempty"`

	ProjectRouting *string `json:"project_routing,omitempty"`

	Source types.ScriptSource `json:"source,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
