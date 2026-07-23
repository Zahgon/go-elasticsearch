package types

type TemplatesRecord struct {
	ComposedOf *string `json:"composed_of,omitempty"`

	IndexPatterns *string `json:"index_patterns,omitempty"`

	Name *string `json:"name,omitempty"`

	Order *string `json:"order,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (s *TemplatesRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTemplatesRecord() *TemplatesRecord { _ = "STUB: not implemented"; return nil }
