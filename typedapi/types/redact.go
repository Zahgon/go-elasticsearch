package types

type Redact struct {
	IsRedacted_ bool `json:"_is_redacted"`
}

func (s *Redact) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRedact() *Redact { _ = "STUB: not implemented"; return nil }
