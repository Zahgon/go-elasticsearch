package types

type DataStreamWithOptions struct {
	Name    string             `json:"name"`
	Options *DataStreamOptions `json:"options,omitempty"`
}

func (s *DataStreamWithOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamWithOptions() *DataStreamWithOptions { _ = "STUB: not implemented"; return nil }
