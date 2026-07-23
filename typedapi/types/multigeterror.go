package types

type MultiGetError struct {
	Error  ErrorCause `json:"error"`
	Id_    string     `json:"_id"`
	Index_ string     `json:"_index"`
}

func (s *MultiGetError) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMultiGetError() *MultiGetError { _ = "STUB: not implemented"; return nil }
