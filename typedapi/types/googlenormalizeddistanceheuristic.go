package types

type GoogleNormalizedDistanceHeuristic struct {
	BackgroundIsSuperset *bool `json:"background_is_superset,omitempty"`
}

func (s *GoogleNormalizedDistanceHeuristic) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGoogleNormalizedDistanceHeuristic() *GoogleNormalizedDistanceHeuristic {
	_ = "STUB: not implemented"
	return nil
}

type GoogleNormalizedDistanceHeuristicVariant interface {
	GoogleNormalizedDistanceHeuristicCaster() *GoogleNormalizedDistanceHeuristic
}

func (s *GoogleNormalizedDistanceHeuristic) GoogleNormalizedDistanceHeuristicCaster() *GoogleNormalizedDistanceHeuristic {
	_ = "STUB: not implemented"
	return nil
}
