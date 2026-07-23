package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _logstashPipeline struct {
	v *types.LogstashPipeline
}

func NewLogstashPipeline(description string, pipeline string, pipelinemetadata types.PipelineMetadataVariant, pipelinesettings types.PipelineSettingsVariant, username string) *_logstashPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_logstashPipeline) Description(description string) *_logstashPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_logstashPipeline) LastModified(datetime types.DateTimeVariant) *_logstashPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_logstashPipeline) Pipeline(pipeline string) *_logstashPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_logstashPipeline) PipelineMetadata(pipelinemetadata types.PipelineMetadataVariant) *_logstashPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_logstashPipeline) PipelineSettings(pipelinesettings types.PipelineSettingsVariant) *_logstashPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_logstashPipeline) Username(username string) *_logstashPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (s *_logstashPipeline) LogstashPipelineCaster() *types.LogstashPipeline {
	_ = "STUB: not implemented"
	return nil
}
