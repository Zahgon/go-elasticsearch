package types

type DataStreamTimestampField struct {
	Name string `json:"name"`
}

func (s *DataStreamTimestampField) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamTimestampField() *DataStreamTimestampField { _ = "STUB: not implemented"; return nil }
