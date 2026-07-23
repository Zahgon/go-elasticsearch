package types

type Ingest struct {
	Pipeline  *string  `json:"pipeline,omitempty"`
	Redact_   *Redact  `json:"_redact,omitempty"`
	Timestamp DateTime `json:"timestamp"`
}

func (s *Ingest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIngest() *Ingest { _ = "STUB: not implemented"; return nil }
