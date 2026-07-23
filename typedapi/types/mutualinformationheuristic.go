package types

type MutualInformationHeuristic struct {
	BackgroundIsSuperset *bool `json:"background_is_superset,omitempty"`

	IncludeNegatives *bool `json:"include_negatives,omitempty"`
}

func (s *MutualInformationHeuristic) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMutualInformationHeuristic() *MutualInformationHeuristic {
	_ = "STUB: not implemented"
	return nil
}

type MutualInformationHeuristicVariant interface {
	MutualInformationHeuristicCaster() *MutualInformationHeuristic
}

func (s *MutualInformationHeuristic) MutualInformationHeuristicCaster() *MutualInformationHeuristic {
	_ = "STUB: not implemented"
	return nil
}
