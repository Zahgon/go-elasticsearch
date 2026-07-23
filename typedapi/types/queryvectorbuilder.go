package types

type QueryVectorBuilder struct {
	Embedding *Embedding `json:"embedding,omitempty"`

	Lookup        *LookupQueryVectorBuilder `json:"lookup,omitempty"`
	TextEmbedding *TextEmbedding            `json:"text_embedding,omitempty"`
}

func NewQueryVectorBuilder() *QueryVectorBuilder { _ = "STUB: not implemented"; return nil }

type QueryVectorBuilderVariant interface {
	QueryVectorBuilderCaster() *QueryVectorBuilder
}

func (s *QueryVectorBuilder) QueryVectorBuilderCaster() *QueryVectorBuilder {
	_ = "STUB: not implemented"
	return nil
}
