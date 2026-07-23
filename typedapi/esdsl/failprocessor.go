package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _failProcessor struct {
	v *types.FailProcessor
}

func NewFailProcessor(message string) *_failProcessor { _ = "STUB: not implemented"; return nil }

func (s *_failProcessor) Message(message string) *_failProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_failProcessor) Description(description string) *_failProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_failProcessor) If(if_ types.ScriptVariant) *_failProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_failProcessor) IgnoreFailure(ignorefailure bool) *_failProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_failProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_failProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_failProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_failProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_failProcessor) Tag(tag string) *_failProcessor { _ = "STUB: not implemented"; return nil }

func (s *_failProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_failProcessor) FailProcessorCaster() *types.FailProcessor {
	_ = "STUB: not implemented"
	return nil
}
