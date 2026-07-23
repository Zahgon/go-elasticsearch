package types

type StatusInProgress struct {
	Index             string `json:"index"`
	ReindexedDocCount int64  `json:"reindexed_doc_count"`
	TotalDocCount     int64  `json:"total_doc_count"`
}

func (s *StatusInProgress) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewStatusInProgress() *StatusInProgress { _ = "STUB: not implemented"; return nil }
