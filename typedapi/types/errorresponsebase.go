package types

type ErrorResponseBase struct {
	Error  ErrorCause `json:"error"`
	Status int        `json:"status"`
}

func (s *ErrorResponseBase) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewErrorResponseBase() *ErrorResponseBase { _ = "STUB: not implemented"; return nil }
