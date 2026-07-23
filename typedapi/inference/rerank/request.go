package rerank

import (
	"encoding/json"
)

type Request struct {
	Input []string `json:"input"`

	Query string `json:"query"`

	ReturnDocuments *bool `json:"return_documents,omitempty"`

	TaskSettings json.RawMessage `json:"task_settings,omitempty"`

	TopN *int `json:"top_n,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
