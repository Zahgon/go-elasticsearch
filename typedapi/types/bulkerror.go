package types

type BulkError struct {
	Count int `json:"count"`

	Details map[string]ErrorCause `json:"details"`
}

func (s *BulkError) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewBulkError() *BulkError { _ = "STUB: not implemented"; return nil }
