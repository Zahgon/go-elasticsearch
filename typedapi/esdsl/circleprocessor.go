package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shapetype"
)

type _circleProcessor struct {
	v *types.CircleProcessor
}

func NewCircleProcessor(errordistance types.Float64, shapetype shapetype.ShapeType) *_circleProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_circleProcessor) ErrorDistance(errordistance types.Float64) *_circleProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_circleProcessor) Field(field string) *_circleProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_circleProcessor) IgnoreMissing(ignoremissing bool) *_circleProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_circleProcessor) ShapeType(shapetype shapetype.ShapeType) *_circleProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_circleProcessor) TargetField(field string) *_circleProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_circleProcessor) Description(description string) *_circleProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_circleProcessor) If(if_ types.ScriptVariant) *_circleProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_circleProcessor) IgnoreFailure(ignorefailure bool) *_circleProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_circleProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_circleProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_circleProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_circleProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_circleProcessor) Tag(tag string) *_circleProcessor { _ = "STUB: not implemented"; return nil }

func (s *_circleProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_circleProcessor) CircleProcessorCaster() *types.CircleProcessor {
	_ = "STUB: not implemented"
	return nil
}
