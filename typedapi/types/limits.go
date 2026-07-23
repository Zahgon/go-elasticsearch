package types

type Limits struct {
	EffectiveMaxModelMemoryLimit ByteSize `json:"effective_max_model_memory_limit,omitempty"`
	MaxModelMemoryLimit          ByteSize `json:"max_model_memory_limit,omitempty"`
	MaxSingleMlNodeProcessors    *int     `json:"max_single_ml_node_processors,omitempty"`
	TotalMlMemory                ByteSize `json:"total_ml_memory"`
	TotalMlProcessors            *int     `json:"total_ml_processors,omitempty"`
}

func (s *Limits) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewLimits() *Limits { _ = "STUB: not implemented"; return nil }
