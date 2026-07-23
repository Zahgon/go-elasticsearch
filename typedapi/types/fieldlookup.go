package types

type FieldLookup struct {
	Id string `json:"id"`

	Index *string `json:"index,omitempty"`

	Path *string `json:"path,omitempty"`

	Routing *string `json:"routing,omitempty"`
}

func (s *FieldLookup) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldLookup() *FieldLookup { _ = "STUB: not implemented"; return nil }

type FieldLookupVariant interface {
	FieldLookupCaster() *FieldLookup
}

func (s *FieldLookup) FieldLookupCaster() *FieldLookup { _ = "STUB: not implemented"; return nil }
