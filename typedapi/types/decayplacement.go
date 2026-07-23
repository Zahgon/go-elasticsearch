package types

import (
	"encoding/json"
)

type DecayPlacement struct {
	Decay *Float64 `json:"decay,omitempty"`

	Offset json.RawMessage `json:"offset,omitempty"`

	Origin json.RawMessage `json:"origin,omitempty"`

	Scale json.RawMessage `json:"scale,omitempty"`
}

func (s *DecayPlacement) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDecayPlacement() *DecayPlacement { _ = "STUB: not implemented"; return nil }

type DecayPlacementVariant interface {
	DecayPlacementCaster() *DecayPlacement
}

func (s *DecayPlacement) DecayPlacementCaster() *DecayPlacement {
	_ = "STUB: not implemented"
	return nil
}
