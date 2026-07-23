package types

type DataStreamFailureStore struct {
	Enabled *bool `json:"enabled,omitempty"`

	Lifecycle *FailureStoreLifecycle `json:"lifecycle,omitempty"`
}

func (s *DataStreamFailureStore) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamFailureStore() *DataStreamFailureStore { _ = "STUB: not implemented"; return nil }

type DataStreamFailureStoreVariant interface {
	DataStreamFailureStoreCaster() *DataStreamFailureStore
}

func (s *DataStreamFailureStore) DataStreamFailureStoreCaster() *DataStreamFailureStore {
	_ = "STUB: not implemented"
	return nil
}
