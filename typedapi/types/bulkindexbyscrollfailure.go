package types

type BulkIndexByScrollFailure struct {
	Cause  ErrorCause `json:"cause"`
	Id     string     `json:"id"`
	Index  string     `json:"index"`
	Status int        `json:"status"`
}

func (s *BulkIndexByScrollFailure) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewBulkIndexByScrollFailure() *BulkIndexByScrollFailure { _ = "STUB: not implemented"; return nil }
