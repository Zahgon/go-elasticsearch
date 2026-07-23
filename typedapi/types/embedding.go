package types

type Embedding struct {
	InferenceId *string           `json:"inference_id,omitempty"`
	Input       KnnEmbeddingInput `json:"input"`
	Timeout     Duration          `json:"timeout,omitempty"`
}

func (s *Embedding) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEmbedding() *Embedding { _ = "STUB: not implemented"; return nil }

type EmbeddingVariant interface {
	EmbeddingCaster() *Embedding
}

func (s *Embedding) EmbeddingCaster() *Embedding { _ = "STUB: not implemented"; return nil }
