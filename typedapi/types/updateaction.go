package types

import (
	"encoding/json"
)

type UpdateAction struct {
	DetectNoop *bool `json:"detect_noop,omitempty"`

	Doc json.RawMessage `json:"doc,omitempty"`

	DocAsUpsert *bool `json:"doc_as_upsert,omitempty"`

	Script *Script `json:"script,omitempty"`

	ScriptedUpsert *bool `json:"scripted_upsert,omitempty"`

	Source_ SourceConfig `json:"_source,omitempty"`

	Upsert json.RawMessage `json:"upsert,omitempty"`
}

func (s *UpdateAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewUpdateAction() *UpdateAction { _ = "STUB: not implemented"; return nil }

type UpdateActionVariant interface {
	UpdateActionCaster() *UpdateAction
}

func (s *UpdateAction) UpdateActionCaster() *UpdateAction { _ = "STUB: not implemented"; return nil }
