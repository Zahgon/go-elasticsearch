package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type TokenCountProperty struct {
	Analyzer                 *string                        `json:"analyzer,omitempty"`
	Boost                    *Float64                       `json:"boost,omitempty"`
	CopyTo                   []string                       `json:"copy_to,omitempty"`
	DocValues                *bool                          `json:"doc_values,omitempty"`
	Dynamic                  *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EnablePositionIncrements *bool                          `json:"enable_position_increments,omitempty"`
	Fields                   map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove              *int                           `json:"ignore_above,omitempty"`
	Index                    *bool                          `json:"index,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	NullValue           *Float64                                         `json:"null_value,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *TokenCountProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s TokenCountProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTokenCountProperty() *TokenCountProperty { _ = "STUB: not implemented"; return nil }

type TokenCountPropertyVariant interface {
	TokenCountPropertyCaster() *TokenCountProperty
}

func (s *TokenCountProperty) TokenCountPropertyCaster() *TokenCountProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *TokenCountProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
