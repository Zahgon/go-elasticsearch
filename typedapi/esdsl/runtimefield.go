package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/runtimefieldtype"
)

type _runtimeField struct {
	v *types.RuntimeField
}

func NewRuntimeField(type_ runtimefieldtype.RuntimeFieldType) *_runtimeField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeField) FetchFields(fetchfields ...types.RuntimeFieldFetchFieldsVariant) *_runtimeField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeField) FetchFieldsValues(fetchfieldsvalues []types.RuntimeFieldFetchFields) *_runtimeField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeField) Fields(fields map[string]types.CompositeSubField) *_runtimeField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeField) AddField(key string, value types.CompositeSubFieldVariant) *_runtimeField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeField) Format(format string) *_runtimeField { _ = "STUB: not implemented"; return nil }

func (s *_runtimeField) InputField(field string) *_runtimeField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeField) OnScriptError(onscripterror onscripterror.OnScriptError) *_runtimeField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeField) Script(script types.ScriptVariant) *_runtimeField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeField) TargetField(field string) *_runtimeField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeField) TargetIndex(indexname string) *_runtimeField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeField) Type(type_ runtimefieldtype.RuntimeFieldType) *_runtimeField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeField) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_runtimeField) RuntimeFieldCaster() *types.RuntimeField {
	_ = "STUB: not implemented"
	return nil
}
