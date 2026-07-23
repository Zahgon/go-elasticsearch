package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _htmlStripProcessor struct {
	v *types.HtmlStripProcessor
}

func NewHtmlStripProcessor() *_htmlStripProcessor { _ = "STUB: not implemented"; return nil }

func (s *_htmlStripProcessor) Field(field string) *_htmlStripProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_htmlStripProcessor) IgnoreMissing(ignoremissing bool) *_htmlStripProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_htmlStripProcessor) TargetField(field string) *_htmlStripProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_htmlStripProcessor) Description(description string) *_htmlStripProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_htmlStripProcessor) If(if_ types.ScriptVariant) *_htmlStripProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_htmlStripProcessor) IgnoreFailure(ignorefailure bool) *_htmlStripProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_htmlStripProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_htmlStripProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_htmlStripProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_htmlStripProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_htmlStripProcessor) Tag(tag string) *_htmlStripProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_htmlStripProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_htmlStripProcessor) HtmlStripProcessorCaster() *types.HtmlStripProcessor {
	_ = "STUB: not implemented"
	return nil
}
