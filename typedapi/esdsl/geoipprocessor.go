package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoIpProcessor struct {
	v *types.GeoIpProcessor
}

func NewGeoIpProcessor() *_geoIpProcessor { _ = "STUB: not implemented"; return nil }

func (s *_geoIpProcessor) DatabaseFile(databasefile string) *_geoIpProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoIpProcessor) DownloadDatabaseOnPipelineCreation(downloaddatabaseonpipelinecreation bool) *_geoIpProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoIpProcessor) Field(field string) *_geoIpProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoIpProcessor) FirstOnly(firstonly bool) *_geoIpProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoIpProcessor) IgnoreMissing(ignoremissing bool) *_geoIpProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoIpProcessor) Properties(properties ...string) *_geoIpProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoIpProcessor) TargetField(field string) *_geoIpProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoIpProcessor) Description(description string) *_geoIpProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoIpProcessor) If(if_ types.ScriptVariant) *_geoIpProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoIpProcessor) IgnoreFailure(ignorefailure bool) *_geoIpProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoIpProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_geoIpProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoIpProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_geoIpProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoIpProcessor) Tag(tag string) *_geoIpProcessor { _ = "STUB: not implemented"; return nil }

func (s *_geoIpProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoIpProcessor) GeoIpProcessorCaster() *types.GeoIpProcessor {
	_ = "STUB: not implemented"
	return nil
}
