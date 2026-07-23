package infertrainedmodel

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Docs []map[string]json.RawMessage `json:"docs"`

	InferenceConfig *types.InferenceConfigUpdateContainer `json:"inference_config,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
