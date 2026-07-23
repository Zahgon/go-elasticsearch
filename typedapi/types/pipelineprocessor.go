package types

type PipelineProcessor struct {
	Description *string `json:"description,omitempty"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissingPipeline *bool `json:"ignore_missing_pipeline,omitempty"`

	Name string `json:"name"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Tag *string `json:"tag,omitempty"`
}

func (s *PipelineProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPipelineProcessor() *PipelineProcessor { _ = "STUB: not implemented"; return nil }

type PipelineProcessorVariant interface {
	PipelineProcessorCaster() *PipelineProcessor
}

func (s *PipelineProcessor) PipelineProcessorCaster() *PipelineProcessor {
	_ = "STUB: not implemented"
	return nil
}
