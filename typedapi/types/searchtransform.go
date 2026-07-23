package types

type SearchTransform struct {
	Request SearchInputRequestDefinition `json:"request"`
	Timeout Duration                     `json:"timeout"`
}

func (s *SearchTransform) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSearchTransform() *SearchTransform { _ = "STUB: not implemented"; return nil }

type SearchTransformVariant interface {
	SearchTransformCaster() *SearchTransform
}

func (s *SearchTransform) SearchTransformCaster() *SearchTransform {
	_ = "STUB: not implemented"
	return nil
}
