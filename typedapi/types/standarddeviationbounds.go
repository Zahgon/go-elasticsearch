package types

type StandardDeviationBounds struct {
	Lower           *Float64 `json:"lower,omitempty"`
	LowerPopulation *Float64 `json:"lower_population,omitempty"`
	LowerSampling   *Float64 `json:"lower_sampling,omitempty"`
	Upper           *Float64 `json:"upper,omitempty"`
	UpperPopulation *Float64 `json:"upper_population,omitempty"`
	UpperSampling   *Float64 `json:"upper_sampling,omitempty"`
}

func (s *StandardDeviationBounds) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewStandardDeviationBounds() *StandardDeviationBounds { _ = "STUB: not implemented"; return nil }
