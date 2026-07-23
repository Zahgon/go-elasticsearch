package types

type KeyedProcessor struct {
	Stats *Processor `json:"stats,omitempty"`
	Type  *string    `json:"type,omitempty"`
}

func (s *KeyedProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewKeyedProcessor() *KeyedProcessor { _ = "STUB: not implemented"; return nil }
