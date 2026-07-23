package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataStreamFailureStore struct {
	v *types.DataStreamFailureStore
}

func NewDataStreamFailureStore() *_dataStreamFailureStore { _ = "STUB: not implemented"; return nil }

func (s *_dataStreamFailureStore) Enabled(enabled bool) *_dataStreamFailureStore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamFailureStore) Lifecycle(lifecycle types.FailureStoreLifecycleVariant) *_dataStreamFailureStore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamFailureStore) DataStreamFailureStoreCaster() *types.DataStreamFailureStore {
	_ = "STUB: not implemented"
	return nil
}
