package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationalternate"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationcasefirst"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationdecomposition"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationstrength"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type IcuCollationProperty struct {
	Alternate              *icucollationalternate.IcuCollationAlternate         `json:"alternate,omitempty"`
	CaseFirst              *icucollationcasefirst.IcuCollationCaseFirst         `json:"case_first,omitempty"`
	CaseLevel              *bool                                                `json:"case_level,omitempty"`
	CopyTo                 []string                                             `json:"copy_to,omitempty"`
	Country                *string                                              `json:"country,omitempty"`
	Decomposition          *icucollationdecomposition.IcuCollationDecomposition `json:"decomposition,omitempty"`
	DocValues              *bool                                                `json:"doc_values,omitempty"`
	Dynamic                *dynamicmapping.DynamicMapping                       `json:"dynamic,omitempty"`
	Fields                 map[string]Property                                  `json:"fields,omitempty"`
	HiraganaQuaternaryMode *bool                                                `json:"hiragana_quaternary_mode,omitempty"`
	IgnoreAbove            *int                                                 `json:"ignore_above,omitempty"`

	Index        *bool                      `json:"index,omitempty"`
	IndexOptions *indexoptions.IndexOptions `json:"index_options,omitempty"`
	Language     *string                    `json:"language,omitempty"`

	Meta  map[string]string `json:"meta,omitempty"`
	Norms *bool             `json:"norms,omitempty"`

	NullValue           *string                                          `json:"null_value,omitempty"`
	Numeric             *bool                                            `json:"numeric,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Rules               *string                                          `json:"rules,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	Strength            *icucollationstrength.IcuCollationStrength       `json:"strength,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
	VariableTop         *string                                          `json:"variable_top,omitempty"`
	Variant             *string                                          `json:"variant,omitempty"`
}

func (s *IcuCollationProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s IcuCollationProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewIcuCollationProperty() *IcuCollationProperty { _ = "STUB: not implemented"; return nil }

type IcuCollationPropertyVariant interface {
	IcuCollationPropertyCaster() *IcuCollationProperty
}

func (s *IcuCollationProperty) IcuCollationPropertyCaster() *IcuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *IcuCollationProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
