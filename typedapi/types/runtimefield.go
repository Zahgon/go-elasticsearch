package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/runtimefieldtype"
)

type RuntimeField struct {
	FetchFields []RuntimeFieldFetchFields `json:"fetch_fields,omitempty"`

	Fields map[string]CompositeSubField `json:"fields,omitempty"`

	Format *string `json:"format,omitempty"`

	InputField    *string                      `json:"input_field,omitempty"`
	OnScriptError *onscripterror.OnScriptError `json:"on_script_error,omitempty"`

	Script *Script `json:"script,omitempty"`

	TargetField *string `json:"target_field,omitempty"`

	TargetIndex *string `json:"target_index,omitempty"`

	Type runtimefieldtype.RuntimeFieldType `json:"type"`
}

func (s *RuntimeField) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRuntimeField() *RuntimeField { _ = "STUB: not implemented"; return nil }

type RuntimeFieldVariant interface {
	RuntimeFieldCaster() *RuntimeField
}

func (s *RuntimeField) RuntimeFieldCaster() *RuntimeField { _ = "STUB: not implemented"; return nil }
