package types

type DataStreamTimestamp struct {
	Enabled bool `json:"enabled"`
}

func (s *DataStreamTimestamp) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamTimestamp() *DataStreamTimestamp { _ = "STUB: not implemented"; return nil }

type DataStreamTimestampVariant interface {
	DataStreamTimestampCaster() *DataStreamTimestamp
}

func (s *DataStreamTimestamp) DataStreamTimestampCaster() *DataStreamTimestamp {
	_ = "STUB: not implemented"
	return nil
}
