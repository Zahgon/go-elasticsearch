package rendersearchtemplate

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	File *string `json:"file,omitempty"`

	Id *string `json:"id,omitempty"`

	Params map[string]json.RawMessage `json:"params,omitempty"`

	Source types.ScriptSource `json:"source,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
