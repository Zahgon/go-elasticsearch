package types

type MultiTermLookup struct {
	Field *string `json:"field,omitempty"`

	Missing Missing `json:"missing,omitempty"`

	Script *Script `json:"script,omitempty"`
}

func (s *MultiTermLookup) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMultiTermLookup() *MultiTermLookup { _ = "STUB: not implemented"; return nil }

type MultiTermLookupVariant interface {
	MultiTermLookupCaster() *MultiTermLookup
}

func (s *MultiTermLookup) MultiTermLookupCaster() *MultiTermLookup {
	_ = "STUB: not implemented"
	return nil
}
