package types

type PercentageScoreHeuristic struct {
}

func NewPercentageScoreHeuristic() *PercentageScoreHeuristic { _ = "STUB: not implemented"; return nil }

type PercentageScoreHeuristicVariant interface {
	PercentageScoreHeuristicCaster() *PercentageScoreHeuristic
}

func (s *PercentageScoreHeuristic) PercentageScoreHeuristicCaster() *PercentageScoreHeuristic {
	_ = "STUB: not implemented"
	return nil
}
