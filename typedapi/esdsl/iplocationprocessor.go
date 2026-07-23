package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ipLocationProcessor struct {
	v *types.IpLocationProcessor
}

func NewIpLocationProcessor() *_ipLocationProcessor { _ = "STUB: not implemented"; return nil }

func (s *_ipLocationProcessor) DatabaseFile(databasefile string) *_ipLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) DownloadDatabaseOnPipelineCreation(downloaddatabaseonpipelinecreation bool) *_ipLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) Field(field string) *_ipLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) FirstOnly(firstonly bool) *_ipLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) IgnoreMissing(ignoremissing bool) *_ipLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) Properties(properties ...string) *_ipLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) TargetField(field string) *_ipLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) Description(description string) *_ipLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) If(if_ types.ScriptVariant) *_ipLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) IgnoreFailure(ignorefailure bool) *_ipLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_ipLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_ipLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) Tag(tag string) *_ipLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ipLocationProcessor) IpLocationProcessorCaster() *types.IpLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}
