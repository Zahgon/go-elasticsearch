package types

type TermVector struct {
	FieldStatistics *FieldStatistics `json:"field_statistics,omitempty"`
	Terms           map[string]Term  `json:"terms"`
}

func NewTermVector() *TermVector { _ = "STUB: not implemented"; return nil }
