package types

type ClusterComponentTemplate struct {
	ComponentTemplate ComponentTemplateNodeWithRollover `json:"component_template"`
	Name              string                            `json:"name"`
}

func (s *ClusterComponentTemplate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClusterComponentTemplate() *ClusterComponentTemplate { _ = "STUB: not implemented"; return nil }
