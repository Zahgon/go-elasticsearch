package types

type IpLocationProcessor struct {
	DatabaseFile *string `json:"database_file,omitempty"`

	Description *string `json:"description,omitempty"`

	DownloadDatabaseOnPipelineCreation *bool `json:"download_database_on_pipeline_creation,omitempty"`

	Field string `json:"field"`

	FirstOnly *bool `json:"first_only,omitempty"`

	If *Script `json:"if,omitempty"`

	IgnoreFailure *bool `json:"ignore_failure,omitempty"`

	IgnoreMissing *bool `json:"ignore_missing,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Properties []string `json:"properties,omitempty"`

	Tag *string `json:"tag,omitempty"`

	TargetField *string `json:"target_field,omitempty"`
}

func (s *IpLocationProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIpLocationProcessor() *IpLocationProcessor { _ = "STUB: not implemented"; return nil }

type IpLocationProcessorVariant interface {
	IpLocationProcessorCaster() *IpLocationProcessor
}

func (s *IpLocationProcessor) IpLocationProcessorCaster() *IpLocationProcessor {
	_ = "STUB: not implemented"
	return nil
}
