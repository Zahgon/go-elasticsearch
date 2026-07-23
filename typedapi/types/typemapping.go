package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/subobjects"
)

type TypeMapping struct {
	AllField             *AllField                      `json:"all_field,omitempty"`
	DataStreamTimestamp_ *DataStreamTimestamp           `json:"_data_stream_timestamp,omitempty"`
	DateDetection        *bool                          `json:"date_detection,omitempty"`
	Dynamic              *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	DynamicDateFormats   []string                       `json:"dynamic_date_formats,omitempty"`
	DynamicTemplates     []map[string]DynamicTemplate   `json:"dynamic_templates,omitempty"`
	Enabled              *bool                          `json:"enabled,omitempty"`
	FieldNames_          *FieldNamesField               `json:"_field_names,omitempty"`
	IndexField           *IndexField                    `json:"index_field,omitempty"`
	Meta_                Metadata                       `json:"_meta,omitempty"`
	NumericDetection     *bool                          `json:"numeric_detection,omitempty"`
	Properties           map[string]Property            `json:"properties,omitempty"`
	Routing_             *RoutingField                  `json:"_routing,omitempty"`
	Runtime              map[string]RuntimeField        `json:"runtime,omitempty"`
	Size_                *SizeField                     `json:"_size,omitempty"`
	Source_              *SourceField                   `json:"_source,omitempty"`
	Subobjects           *subobjects.Subobjects         `json:"subobjects,omitempty"`
}

func (s *TypeMapping) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTypeMapping() *TypeMapping { _ = "STUB: not implemented"; return nil }

type TypeMappingVariant interface {
	TypeMappingCaster() *TypeMapping
}

func (s *TypeMapping) TypeMappingCaster() *TypeMapping { _ = "STUB: not implemented"; return nil }
