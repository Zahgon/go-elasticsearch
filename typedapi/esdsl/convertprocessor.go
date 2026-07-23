package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/converttype"
)

type _convertProcessor struct {
	v *types.ConvertProcessor
}

func NewConvertProcessor(type_ converttype.ConvertType) *_convertProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_convertProcessor) Field(field string) *_convertProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_convertProcessor) IgnoreMissing(ignoremissing bool) *_convertProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_convertProcessor) TargetField(field string) *_convertProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_convertProcessor) Type(type_ converttype.ConvertType) *_convertProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_convertProcessor) Description(description string) *_convertProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_convertProcessor) If(if_ types.ScriptVariant) *_convertProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_convertProcessor) IgnoreFailure(ignorefailure bool) *_convertProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_convertProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_convertProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_convertProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_convertProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_convertProcessor) Tag(tag string) *_convertProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_convertProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_convertProcessor) ConvertProcessorCaster() *types.ConvertProcessor {
	_ = "STUB: not implemented"
	return nil
}
