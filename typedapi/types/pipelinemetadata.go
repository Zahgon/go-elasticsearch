package types

type PipelineMetadata struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

func (s *PipelineMetadata) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPipelineMetadata() *PipelineMetadata { _ = "STUB: not implemented"; return nil }

type PipelineMetadataVariant interface {
	PipelineMetadataCaster() *PipelineMetadata
}

func (s *PipelineMetadata) PipelineMetadataCaster() *PipelineMetadata {
	_ = "STUB: not implemented"
	return nil
}
