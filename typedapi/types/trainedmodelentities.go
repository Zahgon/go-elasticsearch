package types

type TrainedModelEntities struct {
	ClassName        string  `json:"class_name"`
	ClassProbability Float64 `json:"class_probability"`
	EndPos           int     `json:"end_pos"`
	Entity           string  `json:"entity"`
	StartPos         int     `json:"start_pos"`
}

func (s *TrainedModelEntities) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelEntities() *TrainedModelEntities { _ = "STUB: not implemented"; return nil }
