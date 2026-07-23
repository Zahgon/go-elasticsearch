package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geogridtargetformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geogridtiletype"
)

type _geoGridProcessor struct {
	v *types.GeoGridProcessor
}

func NewGeoGridProcessor(field string, tiletype geogridtiletype.GeoGridTileType) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) ChildrenField(field string) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) Field(field string) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) IgnoreMissing(ignoremissing bool) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) NonChildrenField(field string) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) ParentField(field string) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) PrecisionField(field string) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) TargetField(field string) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) TargetFormat(targetformat geogridtargetformat.GeoGridTargetFormat) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) TileType(tiletype geogridtiletype.GeoGridTileType) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) Description(description string) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) If(if_ types.ScriptVariant) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) IgnoreFailure(ignorefailure bool) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) Tag(tag string) *_geoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoGridProcessor) GeoGridProcessorCaster() *types.GeoGridProcessor {
	_ = "STUB: not implemented"
	return nil
}
