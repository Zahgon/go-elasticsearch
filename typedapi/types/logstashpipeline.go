package types

type LogstashPipeline struct {
	Description string `json:"description"`

	LastModified DateTime `json:"last_modified"`

	Pipeline string `json:"pipeline"`

	PipelineMetadata PipelineMetadata `json:"pipeline_metadata"`

	PipelineSettings PipelineSettings `json:"pipeline_settings"`

	Username string `json:"username"`
}

func (s *LogstashPipeline) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewLogstashPipeline() *LogstashPipeline { _ = "STUB: not implemented"; return nil }

type LogstashPipelineVariant interface {
	LogstashPipelineCaster() *LogstashPipeline
}

func (s *LogstashPipeline) LogstashPipelineCaster() *LogstashPipeline {
	_ = "STUB: not implemented"
	return nil
}
