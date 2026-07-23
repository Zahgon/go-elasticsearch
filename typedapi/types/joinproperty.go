package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type JoinProperty struct {
	Dynamic             *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EagerGlobalOrdinals *bool                          `json:"eager_global_ordinals,omitempty"`
	Fields              map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove         *int                           `json:"ignore_above,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Relations           map[string][]string                              `json:"relations,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *JoinProperty) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s JoinProperty) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewJoinProperty() *JoinProperty { _ = "STUB: not implemented"; return nil }

type JoinPropertyVariant interface {
	JoinPropertyCaster() *JoinProperty
}

func (s *JoinProperty) JoinPropertyCaster() *JoinProperty { _ = "STUB: not implemented"; return nil }

func (s *JoinProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
