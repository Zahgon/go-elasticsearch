package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoshaperelation"
)

type _enrichProcessor struct {
	v *types.EnrichProcessor
}

func NewEnrichProcessor(policyname string) *_enrichProcessor { _ = "STUB: not implemented"; return nil }

func (s *_enrichProcessor) Field(field string) *_enrichProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichProcessor) IgnoreMissing(ignoremissing bool) *_enrichProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichProcessor) MaxMatches(maxmatches int) *_enrichProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichProcessor) Override(override bool) *_enrichProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichProcessor) PolicyName(policyname string) *_enrichProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichProcessor) ShapeRelation(shaperelation geoshaperelation.GeoShapeRelation) *_enrichProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichProcessor) TargetField(field string) *_enrichProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichProcessor) Description(description string) *_enrichProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichProcessor) If(if_ types.ScriptVariant) *_enrichProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichProcessor) IgnoreFailure(ignorefailure bool) *_enrichProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_enrichProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_enrichProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichProcessor) Tag(tag string) *_enrichProcessor { _ = "STUB: not implemented"; return nil }

func (s *_enrichProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichProcessor) EnrichProcessorCaster() *types.EnrichProcessor {
	_ = "STUB: not implemented"
	return nil
}
