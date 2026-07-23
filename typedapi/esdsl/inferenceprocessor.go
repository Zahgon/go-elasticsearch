package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _inferenceProcessor struct {
	v *types.InferenceProcessor
}

func NewInferenceProcessor() *_inferenceProcessor { _ = "STUB: not implemented"; return nil }

func (s *_inferenceProcessor) FieldMap(fieldmap map[string]json.RawMessage) *_inferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) AddFieldMap(key string, value json.RawMessage) *_inferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) IgnoreMissing(ignoremissing bool) *_inferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) InferenceConfig(inferenceconfig types.InferenceConfigVariant) *_inferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) InputOutput(inputoutputs ...types.InputConfigVariant) *_inferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) ModelId(id string) *_inferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) TargetField(field string) *_inferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) Description(description string) *_inferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) If(if_ types.ScriptVariant) *_inferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) IgnoreFailure(ignorefailure bool) *_inferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_inferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_inferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) Tag(tag string) *_inferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceProcessor) InferenceProcessorCaster() *types.InferenceProcessor {
	_ = "STUB: not implemented"
	return nil
}
