package types

type PathHierarchyTokenizer struct {
	BufferSize  Stringifiedinteger `json:"buffer_size,omitempty"`
	Delimiter   *string            `json:"delimiter,omitempty"`
	Replacement *string            `json:"replacement,omitempty"`
	Reverse     Stringifiedboolean `json:"reverse,omitempty"`
	Skip        Stringifiedinteger `json:"skip,omitempty"`
	Type        string             `json:"type,omitempty"`
	Version     *string            `json:"version,omitempty"`
}

func (s *PathHierarchyTokenizer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s PathHierarchyTokenizer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPathHierarchyTokenizer() *PathHierarchyTokenizer { _ = "STUB: not implemented"; return nil }

type PathHierarchyTokenizerVariant interface {
	PathHierarchyTokenizerCaster() *PathHierarchyTokenizer
}

func (s *PathHierarchyTokenizer) PathHierarchyTokenizerCaster() *PathHierarchyTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *PathHierarchyTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
