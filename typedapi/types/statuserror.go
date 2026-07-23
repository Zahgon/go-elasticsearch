package types

type StatusError struct {
	Index   string `json:"index"`
	Message string `json:"message"`
}

func (s *StatusError) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewStatusError() *StatusError { _ = "STUB: not implemented"; return nil }
