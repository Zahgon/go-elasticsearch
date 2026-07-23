package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type DateProperty struct {
	Boost           *Float64                       `json:"boost,omitempty"`
	CopyTo          []string                       `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fielddata       *NumericFielddata              `json:"fielddata,omitempty"`
	Fields          map[string]Property            `json:"fields,omitempty"`
	Format          *string                        `json:"format,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	Index           *bool                          `json:"index,omitempty"`
	Locale          *string                        `json:"locale,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	NullValue           DateTime                                         `json:"null_value,omitempty"`
	OnScriptError       *onscripterror.OnScriptError                     `json:"on_script_error,omitempty"`
	PrecisionStep       *int                                             `json:"precision_step,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Script              *Script                                          `json:"script,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *DateProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s DateProperty) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDateProperty() *DateProperty { _ = "STUB: not implemented"; return nil }

type DatePropertyVariant interface {
	DatePropertyCaster() *DateProperty
}

func (s *DateProperty) DatePropertyCaster() *DateProperty { _ = "STUB: not implemented"; return nil }

func (s *DateProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
