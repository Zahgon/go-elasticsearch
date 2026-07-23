package test

import (
	"encoding/json"
)

type Request struct {
	MatchCriteria map[string]json.RawMessage `json:"match_criteria"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
