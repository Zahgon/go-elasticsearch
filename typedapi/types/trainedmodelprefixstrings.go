package types

type TrainedModelPrefixStrings struct {
	Ingest *string `json:"ingest,omitempty"`

	Search *string `json:"search,omitempty"`
}

func (s *TrainedModelPrefixStrings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelPrefixStrings() *TrainedModelPrefixStrings {
	_ = "STUB: not implemented"
	return nil
}

type TrainedModelPrefixStringsVariant interface {
	TrainedModelPrefixStringsCaster() *TrainedModelPrefixStrings
}

func (s *TrainedModelPrefixStrings) TrainedModelPrefixStringsCaster() *TrainedModelPrefixStrings {
	_ = "STUB: not implemented"
	return nil
}
