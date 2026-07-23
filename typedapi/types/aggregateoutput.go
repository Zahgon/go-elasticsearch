package types

type AggregateOutput struct {
	Exponent           *Weights `json:"exponent,omitempty"`
	LogisticRegression *Weights `json:"logistic_regression,omitempty"`
	WeightedMode       *Weights `json:"weighted_mode,omitempty"`
	WeightedSum        *Weights `json:"weighted_sum,omitempty"`
}

func NewAggregateOutput() *AggregateOutput { _ = "STUB: not implemented"; return nil }

type AggregateOutputVariant interface {
	AggregateOutputCaster() *AggregateOutput
}

func (s *AggregateOutput) AggregateOutputCaster() *AggregateOutput {
	_ = "STUB: not implemented"
	return nil
}
