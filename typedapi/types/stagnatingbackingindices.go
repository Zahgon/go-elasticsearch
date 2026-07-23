package types

type StagnatingBackingIndices struct {
	FirstOccurrenceTimestamp int64  `json:"first_occurrence_timestamp"`
	IndexName                string `json:"index_name"`
	RetryCount               int    `json:"retry_count"`
}

func (s *StagnatingBackingIndices) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewStagnatingBackingIndices() *StagnatingBackingIndices { _ = "STUB: not implemented"; return nil }
