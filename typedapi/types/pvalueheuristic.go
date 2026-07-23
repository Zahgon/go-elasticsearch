package types

type PValueHeuristic struct {
	BackgroundIsSuperset *bool `json:"background_is_superset,omitempty"`

	NormalizeAbove *int64 `json:"normalize_above,omitempty"`
}

func (s *PValueHeuristic) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPValueHeuristic() *PValueHeuristic { _ = "STUB: not implemented"; return nil }

type PValueHeuristicVariant interface {
	PValueHeuristicCaster() *PValueHeuristic
}

func (s *PValueHeuristic) PValueHeuristicCaster() *PValueHeuristic {
	_ = "STUB: not implemented"
	return nil
}
