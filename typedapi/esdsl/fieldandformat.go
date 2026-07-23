package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fieldAndFormat struct {
	v *types.FieldAndFormat
}

func NewFieldAndFormat() *_fieldAndFormat { _ = "STUB: not implemented"; return nil }

func (s *_fieldAndFormat) Field(field string) *_fieldAndFormat {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAndFormat) Format(format string) *_fieldAndFormat {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAndFormat) IncludeUnmapped(includeunmapped bool) *_fieldAndFormat {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldAndFormat) FieldAndFormatCaster() *types.FieldAndFormat {
	_ = "STUB: not implemented"
	return nil
}
