package types

type RandomScoreFunction struct {
	Field *string `json:"field,omitempty"`
	Seed  *string `json:"seed,omitempty"`
}

func (s *RandomScoreFunction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRandomScoreFunction() *RandomScoreFunction { _ = "STUB: not implemented"; return nil }

type RandomScoreFunctionVariant interface {
	RandomScoreFunctionCaster() *RandomScoreFunction
}

func (s *RandomScoreFunction) RandomScoreFunctionCaster() *RandomScoreFunction {
	_ = "STUB: not implemented"
	return nil
}
