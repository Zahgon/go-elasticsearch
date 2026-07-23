package types

type Audit struct {
	Enabled bool     `json:"enabled"`
	Outputs []string `json:"outputs,omitempty"`
}

func (s *Audit) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAudit() *Audit { _ = "STUB: not implemented"; return nil }
