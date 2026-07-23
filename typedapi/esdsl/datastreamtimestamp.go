package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataStreamTimestamp struct {
	v *types.DataStreamTimestamp
}

func NewDataStreamTimestamp(enabled bool) *_dataStreamTimestamp {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamTimestamp) Enabled(enabled bool) *_dataStreamTimestamp {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamTimestamp) DataStreamTimestampCaster() *types.DataStreamTimestamp {
	_ = "STUB: not implemented"
	return nil
}
