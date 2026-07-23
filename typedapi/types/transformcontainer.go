package types

type TransformContainer struct {
	Chain  []TransformContainer `json:"chain,omitempty"`
	Script *ScriptTransform     `json:"script,omitempty"`
	Search *SearchTransform     `json:"search,omitempty"`
}

func NewTransformContainer() *TransformContainer { _ = "STUB: not implemented"; return nil }

type TransformContainerVariant interface {
	TransformContainerCaster() *TransformContainer
}

func (s *TransformContainer) TransformContainerCaster() *TransformContainer {
	_ = "STUB: not implemented"
	return nil
}
