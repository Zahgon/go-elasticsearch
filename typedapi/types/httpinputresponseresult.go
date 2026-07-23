package types

type HttpInputResponseResult struct {
	Body    string      `json:"body"`
	Headers HttpHeaders `json:"headers"`
	Status  int         `json:"status"`
}

func (s *HttpInputResponseResult) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHttpInputResponseResult() *HttpInputResponseResult { _ = "STUB: not implemented"; return nil }
