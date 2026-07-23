package types

type BucketCorrelationFunction struct {
	CountCorrelation BucketCorrelationFunctionCountCorrelation `json:"count_correlation"`
}

func NewBucketCorrelationFunction() *BucketCorrelationFunction {
	_ = "STUB: not implemented"
	return nil
}

type BucketCorrelationFunctionVariant interface {
	BucketCorrelationFunctionCaster() *BucketCorrelationFunction
}

func (s *BucketCorrelationFunction) BucketCorrelationFunctionCaster() *BucketCorrelationFunction {
	_ = "STUB: not implemented"
	return nil
}
