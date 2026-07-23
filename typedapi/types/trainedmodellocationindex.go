package types

type TrainedModelLocationIndex struct {
	Name string `json:"name"`
}

func (s *TrainedModelLocationIndex) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelLocationIndex() *TrainedModelLocationIndex {
	_ = "STUB: not implemented"
	return nil
}
