package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _bucketCorrelationFunction struct {
	v *types.BucketCorrelationFunction
}

func NewBucketCorrelationFunction(countcorrelation types.BucketCorrelationFunctionCountCorrelationVariant) *_bucketCorrelationFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketCorrelationFunction) CountCorrelation(countcorrelation types.BucketCorrelationFunctionCountCorrelationVariant) *_bucketCorrelationFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketCorrelationFunction) BucketCorrelationFunctionCaster() *types.BucketCorrelationFunction {
	_ = "STUB: not implemented"
	return nil
}
