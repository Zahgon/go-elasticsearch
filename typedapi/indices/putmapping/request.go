package putmapping

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
)

type Request struct {
	DateDetection *bool `json:"date_detection,omitempty"`

	Dynamic *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`

	DynamicDateFormats []string `json:"dynamic_date_formats,omitempty"`

	DynamicTemplates []map[string]types.DynamicTemplate `json:"dynamic_templates,omitempty"`

	FieldNames_ *types.FieldNamesField `json:"_field_names,omitempty"`

	Meta_ types.Metadata `json:"_meta,omitempty"`

	NumericDetection *bool `json:"numeric_detection,omitempty"`

	Properties map[string]types.Property `json:"properties,omitempty"`

	Routing_ *types.RoutingField `json:"_routing,omitempty"`

	Runtime types.RuntimeFields `json:"runtime,omitempty"`

	Source_ *types.SourceField `json:"_source,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
