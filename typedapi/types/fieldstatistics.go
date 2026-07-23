package types

type FieldStatistics struct {
	DocCount   int   `json:"doc_count"`
	SumDocFreq int64 `json:"sum_doc_freq"`
	SumTtf     int64 `json:"sum_ttf"`
}

func (s *FieldStatistics) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldStatistics() *FieldStatistics { _ = "STUB: not implemented"; return nil }
