package types

type SpecifiedDocument struct {
	Id    string  `json:"id"`
	Index *string `json:"index,omitempty"`
}

func (s *SpecifiedDocument) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSpecifiedDocument() *SpecifiedDocument { _ = "STUB: not implemented"; return nil }

type SpecifiedDocumentVariant interface {
	SpecifiedDocumentCaster() *SpecifiedDocument
}

func (s *SpecifiedDocument) SpecifiedDocumentCaster() *SpecifiedDocument {
	_ = "STUB: not implemented"
	return nil
}
