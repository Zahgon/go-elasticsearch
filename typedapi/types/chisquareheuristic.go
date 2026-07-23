package types

type ChiSquareHeuristic struct {
	BackgroundIsSuperset bool `json:"background_is_superset"`

	IncludeNegatives bool `json:"include_negatives"`
}

func (s *ChiSquareHeuristic) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewChiSquareHeuristic() *ChiSquareHeuristic { _ = "STUB: not implemented"; return nil }

type ChiSquareHeuristicVariant interface {
	ChiSquareHeuristicCaster() *ChiSquareHeuristic
}

func (s *ChiSquareHeuristic) ChiSquareHeuristicCaster() *ChiSquareHeuristic {
	_ = "STUB: not implemented"
	return nil
}
