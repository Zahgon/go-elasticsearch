package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fieldaccesspattern"
)

type IngestPipeline struct {
	CreatedDate DateTime `json:"created_date,omitempty"`

	CreatedDateMillis *int64 `json:"created_date_millis,omitempty"`

	Deprecated *bool `json:"deprecated,omitempty"`

	Description *string `json:"description,omitempty"`

	FieldAccessPattern *fieldaccesspattern.FieldAccessPattern `json:"field_access_pattern,omitempty"`

	Meta_ Metadata `json:"_meta,omitempty"`

	ModifiedDate DateTime `json:"modified_date,omitempty"`

	ModifiedDateMillis *int64 `json:"modified_date_millis,omitempty"`

	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`

	Processors []ProcessorContainer `json:"processors,omitempty"`

	Version *int64 `json:"version,omitempty"`
}

func (s *IngestPipeline) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIngestPipeline() *IngestPipeline { _ = "STUB: not implemented"; return nil }

type IngestPipelineVariant interface {
	IngestPipelineCaster() *IngestPipeline
}

func (s *IngestPipeline) IngestPipelineCaster() *IngestPipeline {
	_ = "STUB: not implemented"
	return nil
}
