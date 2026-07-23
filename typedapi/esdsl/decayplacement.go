package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _decayPlacement struct {
	v *types.DecayPlacement
}

func NewDecayPlacement() *_decayPlacement { _ = "STUB: not implemented"; return nil }

func (s *_decayPlacement) Decay(decay types.Float64) *_decayPlacement {
	_ = "STUB: not implemented"
	return nil
}

func (s *_decayPlacement) Offset(offset json.RawMessage) *_decayPlacement {
	_ = "STUB: not implemented"
	return nil
}

func (s *_decayPlacement) Origin(origin json.RawMessage) *_decayPlacement {
	_ = "STUB: not implemented"
	return nil
}

func (s *_decayPlacement) Scale(scale json.RawMessage) *_decayPlacement {
	_ = "STUB: not implemented"
	return nil
}

func (s *_decayPlacement) DecayPlacementCaster() *types.DecayPlacement {
	_ = "STUB: not implemented"
	return nil
}
