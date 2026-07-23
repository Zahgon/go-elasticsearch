package types

type KnnCollectorResult struct {
	Children    []KnnCollectorResult `json:"children,omitempty"`
	Name        string               `json:"name"`
	Reason      string               `json:"reason"`
	Time        Duration             `json:"time,omitempty"`
	TimeInNanos int64                `json:"time_in_nanos"`
}

func (s *KnnCollectorResult) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewKnnCollectorResult() *KnnCollectorResult { _ = "STUB: not implemented"; return nil }
