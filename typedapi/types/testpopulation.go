package types

type TestPopulation struct {
	Field string `json:"field"`

	Filter *Query  `json:"filter,omitempty"`
	Script *Script `json:"script,omitempty"`
}

func (s *TestPopulation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTestPopulation() *TestPopulation { _ = "STUB: not implemented"; return nil }

type TestPopulationVariant interface {
	TestPopulationCaster() *TestPopulation
}

func (s *TestPopulation) TestPopulationCaster() *TestPopulation {
	_ = "STUB: not implemented"
	return nil
}
