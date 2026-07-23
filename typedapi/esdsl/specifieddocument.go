package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _specifiedDocument struct {
	v *types.SpecifiedDocument
}

func NewSpecifiedDocument() *_specifiedDocument { _ = "STUB: not implemented"; return nil }

func (s *_specifiedDocument) Id(id string) *_specifiedDocument {
	_ = "STUB: not implemented"
	return nil
}

func (s *_specifiedDocument) Index(indexname string) *_specifiedDocument {
	_ = "STUB: not implemented"
	return nil
}

func (s *_specifiedDocument) SpecifiedDocumentCaster() *types.SpecifiedDocument {
	_ = "STUB: not implemented"
	return nil
}
