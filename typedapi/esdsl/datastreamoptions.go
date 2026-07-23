package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataStreamOptions struct {
	v *types.DataStreamOptions
}

func NewDataStreamOptions() *_dataStreamOptions { _ = "STUB: not implemented"; return nil }

func (s *_dataStreamOptions) FailureStore(failurestore types.DataStreamFailureStoreVariant) *_dataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamOptions) DataStreamOptionsCaster() *types.DataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}
