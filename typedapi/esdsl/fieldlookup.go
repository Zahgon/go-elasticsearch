package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fieldLookup struct {
	v *types.FieldLookup
}

func NewFieldLookup() *_fieldLookup { _ = "STUB: not implemented"; return nil }

func (s *_fieldLookup) Id(id string) *_fieldLookup { _ = "STUB: not implemented"; return nil }

func (s *_fieldLookup) Index(indexname string) *_fieldLookup { _ = "STUB: not implemented"; return nil }

func (s *_fieldLookup) Path(field string) *_fieldLookup { _ = "STUB: not implemented"; return nil }

func (s *_fieldLookup) Routing(routing string) *_fieldLookup { _ = "STUB: not implemented"; return nil }

func (s *_fieldLookup) FieldLookupCaster() *types.FieldLookup {
	_ = "STUB: not implemented"
	return nil
}
