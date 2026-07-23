package types

type SemanticTextProperty struct {
	ChunkingSettings *ChunkingSettings `json:"chunking_settings,omitempty"`

	Fields map[string]Property `json:"fields,omitempty"`

	IndexOptions *SemanticTextIndexOptions `json:"index_options,omitempty"`

	InferenceId *string           `json:"inference_id,omitempty"`
	Meta        map[string]string `json:"meta,omitempty"`

	SearchInferenceId *string `json:"search_inference_id,omitempty"`
	Type              string  `json:"type,omitempty"`
}

func (s *SemanticTextProperty) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SemanticTextProperty) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSemanticTextProperty() *SemanticTextProperty { _ = "STUB: not implemented"; return nil }

type SemanticTextPropertyVariant interface {
	SemanticTextPropertyCaster() *SemanticTextProperty
}

func (s *SemanticTextProperty) SemanticTextPropertyCaster() *SemanticTextProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *SemanticTextProperty) PropertyCaster() *Property { _ = "STUB: not implemented"; return nil }
