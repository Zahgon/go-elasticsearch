package types

type BucketCorrelationFunctionCountCorrelationIndicator struct {
	DocCount int `json:"doc_count"`

	Expectations []Float64 `json:"expectations"`

	Fractions []Float64 `json:"fractions,omitempty"`
}

func (s *BucketCorrelationFunctionCountCorrelationIndicator) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewBucketCorrelationFunctionCountCorrelationIndicator() *BucketCorrelationFunctionCountCorrelationIndicator {
	_ = "STUB: not implemented"
	return nil
}

type BucketCorrelationFunctionCountCorrelationIndicatorVariant interface {
	BucketCorrelationFunctionCountCorrelationIndicatorCaster() *BucketCorrelationFunctionCountCorrelationIndicator
}

func (s *BucketCorrelationFunctionCountCorrelationIndicator) BucketCorrelationFunctionCountCorrelationIndicatorCaster() *BucketCorrelationFunctionCountCorrelationIndicator {
	_ = "STUB: not implemented"
	return nil
}
