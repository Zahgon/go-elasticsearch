package types

import (
	"encoding/json"
)

type PainlessContextSetup struct {
	Document json.RawMessage `json:"document,omitempty"`

	Index string `json:"index"`

	Query *Query `json:"query,omitempty"`
}

func (s *PainlessContextSetup) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPainlessContextSetup() *PainlessContextSetup { _ = "STUB: not implemented"; return nil }

type PainlessContextSetupVariant interface {
	PainlessContextSetupCaster() *PainlessContextSetup
}

func (s *PainlessContextSetup) PainlessContextSetupCaster() *PainlessContextSetup {
	_ = "STUB: not implemented"
	return nil
}
