package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type PassthroughObjectProperty struct {
	CopyTo      []string                       `json:"copy_to,omitempty"`
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Enabled     *bool                          `json:"enabled,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`

	Meta                map[string]string                                `json:"meta,omitempty"`
	Priority            *int                                             `json:"priority,omitempty"`
	Properties          map[string]Property                              `json:"properties,omitempty"`
	Store               *bool                                            `json:"store,omitempty"`
	SyntheticSourceKeep *syntheticsourcekeepenum.SyntheticSourceKeepEnum `json:"synthetic_source_keep,omitempty"`
	TimeSeriesDimension *bool                                            `json:"time_series_dimension,omitempty"`
	Type                string                                           `json:"type,omitempty"`
}

func (s *PassthroughObjectProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s PassthroughObjectProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPassthroughObjectProperty() *PassthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

type PassthroughObjectPropertyVariant interface {
	PassthroughObjectPropertyCaster() *PassthroughObjectProperty
}

func (s *PassthroughObjectProperty) PassthroughObjectPropertyCaster() *PassthroughObjectProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *PassthroughObjectProperty) PropertyCaster() *Property {
	_ = "STUB: not implemented"
	return nil
}
