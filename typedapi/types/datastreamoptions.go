package types

type DataStreamOptions struct {
	FailureStore *DataStreamFailureStore `json:"failure_store,omitempty"`
}

func NewDataStreamOptions() *DataStreamOptions { _ = "STUB: not implemented"; return nil }

type DataStreamOptionsVariant interface {
	DataStreamOptionsCaster() *DataStreamOptions
}

func (s *DataStreamOptions) DataStreamOptionsCaster() *DataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}
