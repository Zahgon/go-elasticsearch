package types

type BucketCorrelationFunctionCountCorrelation struct {
	Indicator BucketCorrelationFunctionCountCorrelationIndicator `json:"indicator"`
}

func NewBucketCorrelationFunctionCountCorrelation() *BucketCorrelationFunctionCountCorrelation {
	_ = "STUB: not implemented"
	return nil
}

type BucketCorrelationFunctionCountCorrelationVariant interface {
	BucketCorrelationFunctionCountCorrelationCaster() *BucketCorrelationFunctionCountCorrelation
}

func (s *BucketCorrelationFunctionCountCorrelation) BucketCorrelationFunctionCountCorrelationCaster() *BucketCorrelationFunctionCountCorrelation {
	_ = "STUB: not implemented"
	return nil
}
