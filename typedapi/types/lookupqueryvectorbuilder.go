package types

type LookupQueryVectorBuilder struct {
	Id string `json:"id"`

	Index string `json:"index"`

	Path string `json:"path"`

	Routing *string `json:"routing,omitempty"`
}

func (s *LookupQueryVectorBuilder) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewLookupQueryVectorBuilder() *LookupQueryVectorBuilder { _ = "STUB: not implemented"; return nil }

type LookupQueryVectorBuilderVariant interface {
	LookupQueryVectorBuilderCaster() *LookupQueryVectorBuilder
}

func (s *LookupQueryVectorBuilder) LookupQueryVectorBuilderCaster() *LookupQueryVectorBuilder {
	_ = "STUB: not implemented"
	return nil
}
