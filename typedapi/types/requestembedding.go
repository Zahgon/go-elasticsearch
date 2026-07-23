package types

import (
	"encoding/json"
)

type RequestEmbedding struct {
	Input EmbeddingInput `json:"input"`

	InputType *string `json:"input_type,omitempty"`

	TaskSettings json.RawMessage `json:"task_settings,omitempty"`
}

func (s *RequestEmbedding) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRequestEmbedding() *RequestEmbedding { _ = "STUB: not implemented"; return nil }

type RequestEmbeddingVariant interface {
	RequestEmbeddingCaster() *RequestEmbedding
}

func (s *RequestEmbedding) RequestEmbeddingCaster() *RequestEmbedding {
	_ = "STUB: not implemented"
	return nil
}
