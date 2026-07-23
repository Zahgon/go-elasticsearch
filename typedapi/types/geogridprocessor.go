package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geogridtargetformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geogridtiletype"
)

type GeoGridProcessor struct {
	ChildrenField *string `json:"children_field,omitempty"`

	Description *string `json:"description,omitempty"`

	Field string `json:"field"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	NonChildrenField *string `json:"non_children_field,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	ParentField *string `json:"parent_field,omitempty"`

	PrecisionField *string `json:"precision_field,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`

	TargetFormat *geogridtargetformat.GeoGridTargetFormat `json:"target_format,omitempty"`

	TileType geogridtiletype.GeoGridTileType `json:"tile_type"`
}

func (s *GeoGridProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGeoGridProcessor() *GeoGridProcessor { _ = "STUB: not implemented"; return nil }

type GeoGridProcessorVariant interface {
	GeoGridProcessorCaster() *GeoGridProcessor
}

func (s *GeoGridProcessor) GeoGridProcessorCaster() *GeoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}
