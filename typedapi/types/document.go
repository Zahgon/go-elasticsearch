package types

import (
	"encoding/json"
)

type Document struct {
	Id_ *string `json:"_id,omitempty"`

	Index_ *string `json:"_index,omitempty"`

	Source_ json.RawMessage `json:"_source,omitempty"`
}

func (s *Document) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDocument() *Document { _ = "STUB: not implemented"; return nil }

type DocumentVariant interface {
	DocumentCaster() *Document
}

func (s *Document) DocumentCaster() *Document { _ = "STUB: not implemented"; return nil }
