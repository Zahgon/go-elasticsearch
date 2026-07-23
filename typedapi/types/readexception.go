package types

type ReadException struct {
	Exception ErrorCause `json:"exception"`

	FromSeqNo int64 `json:"from_seq_no"`

	Retries int `json:"retries"`
}

func (s *ReadException) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewReadException() *ReadException { _ = "STUB: not implemented"; return nil }
