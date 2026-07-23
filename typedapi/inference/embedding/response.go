package embedding

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Embeddings      []types.DenseEmbeddingResult     `json:"embeddings,omitempty"`
	EmbeddingsBits  []types.DenseEmbeddingByteResult `json:"embeddings_bits,omitempty"`
	EmbeddingsBytes []types.DenseEmbeddingByteResult `json:"embeddings_bytes,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
