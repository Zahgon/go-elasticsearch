package types

type TransformSource struct {
	Index []string `json:"index"`

	ProjectRouting *string `json:"project_routing,omitempty"`

	Query *Query `json:"query,omitempty"`

	RuntimeMappings RuntimeFields `json:"runtime_mappings,omitempty"`
}

func (s *TransformSource) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTransformSource() *TransformSource { _ = "STUB: not implemented"; return nil }

type TransformSourceVariant interface {
	TransformSourceCaster() *TransformSource
}

func (s *TransformSource) TransformSourceCaster() *TransformSource {
	_ = "STUB: not implemented"
	return nil
}
