package types

type DataStreamLifecycleDetails struct {
	StagnatingBackingIndices      []StagnatingBackingIndices `json:"stagnating_backing_indices,omitempty"`
	StagnatingBackingIndicesCount int                        `json:"stagnating_backing_indices_count"`
	TotalBackingIndicesInError    int                        `json:"total_backing_indices_in_error"`
}

func (s *DataStreamLifecycleDetails) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamLifecycleDetails() *DataStreamLifecycleDetails {
	_ = "STUB: not implemented"
	return nil
}
