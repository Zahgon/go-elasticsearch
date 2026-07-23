package types

type TextEmbedding struct {
	ModelId *string `json:"model_id,omitempty"`

	ModelText string `json:"model_text"`
}

func (s *TextEmbedding) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTextEmbedding() *TextEmbedding { _ = "STUB: not implemented"; return nil }

type TextEmbeddingVariant interface {
	TextEmbeddingCaster() *TextEmbedding
}

func (s *TextEmbedding) TextEmbeddingCaster() *TextEmbedding { _ = "STUB: not implemented"; return nil }
