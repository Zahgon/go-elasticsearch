package types

type ExtendedStatsAggregate struct {
	Avg                        *Float64                         `json:"avg,omitempty"`
	AvgAsString                *string                          `json:"avg_as_string,omitempty"`
	Count                      int64                            `json:"count"`
	Max                        *Float64                         `json:"max,omitempty"`
	MaxAsString                *string                          `json:"max_as_string,omitempty"`
	Meta                       Metadata                         `json:"meta,omitempty"`
	Min                        *Float64                         `json:"min,omitempty"`
	MinAsString                *string                          `json:"min_as_string,omitempty"`
	StdDeviation               *Float64                         `json:"std_deviation,omitempty"`
	StdDeviationAsString       *string                          `json:"std_deviation_as_string,omitempty"`
	StdDeviationBounds         *StandardDeviationBounds         `json:"std_deviation_bounds,omitempty"`
	StdDeviationBoundsAsString *StandardDeviationBoundsAsString `json:"std_deviation_bounds_as_string,omitempty"`
	StdDeviationPopulation     *Float64                         `json:"std_deviation_population,omitempty"`
	StdDeviationSampling       *Float64                         `json:"std_deviation_sampling,omitempty"`
	Sum                        Float64                          `json:"sum"`
	SumAsString                *string                          `json:"sum_as_string,omitempty"`
	SumOfSquares               *Float64                         `json:"sum_of_squares,omitempty"`
	SumOfSquaresAsString       *string                          `json:"sum_of_squares_as_string,omitempty"`
	Variance                   *Float64                         `json:"variance,omitempty"`
	VarianceAsString           *string                          `json:"variance_as_string,omitempty"`
	VariancePopulation         *Float64                         `json:"variance_population,omitempty"`
	VariancePopulationAsString *string                          `json:"variance_population_as_string,omitempty"`
	VarianceSampling           *Float64                         `json:"variance_sampling,omitempty"`
	VarianceSamplingAsString   *string                          `json:"variance_sampling_as_string,omitempty"`
}

func (s *ExtendedStatsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewExtendedStatsAggregate() *ExtendedStatsAggregate { _ = "STUB: not implemented"; return nil }
