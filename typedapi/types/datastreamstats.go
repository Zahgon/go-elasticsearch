package types

type DataStreamStats struct {
	BackingIndicesInError int `json:"backing_indices_in_error"`

	BackingIndicesInTotal int `json:"backing_indices_in_total"`

	Name string `json:"name"`
}

func (s *DataStreamStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDataStreamStats() *DataStreamStats { _ = "STUB: not implemented"; return nil }
