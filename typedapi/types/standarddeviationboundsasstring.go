package types

type StandardDeviationBoundsAsString struct {
	Lower           string `json:"lower"`
	LowerPopulation string `json:"lower_population"`
	LowerSampling   string `json:"lower_sampling"`
	Upper           string `json:"upper"`
	UpperPopulation string `json:"upper_population"`
	UpperSampling   string `json:"upper_sampling"`
}

func (s *StandardDeviationBoundsAsString) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewStandardDeviationBoundsAsString() *StandardDeviationBoundsAsString {
	_ = "STUB: not implemented"
	return nil
}
