package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fieldaccesspattern"
)

type _ingestPipeline struct {
	v *types.IngestPipeline
}

func NewIngestPipeline() *_ingestPipeline { _ = "STUB: not implemented"; return nil }

func (s *_ingestPipeline) CreatedDate(datetime types.DateTimeVariant) *_ingestPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipeline) CreatedDateMillis(epochtimeunitmillis int64) *_ingestPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipeline) Deprecated(deprecated bool) *_ingestPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipeline) Description(description string) *_ingestPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipeline) FieldAccessPattern(fieldaccesspattern fieldaccesspattern.FieldAccessPattern) *_ingestPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipeline) Meta_(metadata types.MetadataVariant) *_ingestPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipeline) ModifiedDate(datetime types.DateTimeVariant) *_ingestPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipeline) ModifiedDateMillis(epochtimeunitmillis int64) *_ingestPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipeline) OnFailure(onfailures ...types.ProcessorContainerVariant) *_ingestPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipeline) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_ingestPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipeline) Processors(processors ...types.ProcessorContainerVariant) *_ingestPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipeline) ProcessorsValues(processorsvalues []types.ProcessorContainer) *_ingestPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipeline) Version(versionnumber int64) *_ingestPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ingestPipeline) IngestPipelineCaster() *types.IngestPipeline {
	_ = "STUB: not implemented"
	return nil
}
