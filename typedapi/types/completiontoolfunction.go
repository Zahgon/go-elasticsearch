package types

import (
	"encoding/json"
)

type CompletionToolFunction struct {
	Description *string `json:"description,omitempty"`

	Name string `json:"name"`

	Parameters json.RawMessage `json:"parameters,omitempty"`

	Strict *bool `json:"strict,omitempty"`
}

func (s *CompletionToolFunction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCompletionToolFunction() *CompletionToolFunction { _ = "STUB: not implemented"; return nil }

type CompletionToolFunctionVariant interface {
	CompletionToolFunctionCaster() *CompletionToolFunction
}

func (s *CompletionToolFunction) CompletionToolFunctionCaster() *CompletionToolFunction {
	_ = "STUB: not implemented"
	return nil
}
