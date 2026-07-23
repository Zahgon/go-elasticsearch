package types

type PipelineConfig struct {
	Description *string `json:"description,omitempty"`

	Processors []ProcessorContainer `json:"processors"`

	Version *int64 `json:"version,omitempty"`
}

func (s *PipelineConfig) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPipelineConfig() *PipelineConfig { _ = "STUB: not implemented"; return nil }
