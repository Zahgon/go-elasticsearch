package inference

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Completion         []types.CompletionResult         `json:"completion,omitempty"`
	Embeddings         []types.DenseEmbeddingResult     `json:"embeddings,omitempty"`
	EmbeddingsBits     []types.DenseEmbeddingByteResult `json:"embeddings_bits,omitempty"`
	EmbeddingsBytes    []types.DenseEmbeddingByteResult `json:"embeddings_bytes,omitempty"`
	Rerank             []types.RankedDocument           `json:"rerank,omitempty"`
	SparseEmbedding    []types.SparseEmbeddingResult    `json:"sparse_embedding,omitempty"`
	TextEmbedding      []types.DenseEmbeddingResult     `json:"text_embedding,omitempty"`
	TextEmbeddingBits  []types.DenseEmbeddingByteResult `json:"text_embedding_bits,omitempty"`
	TextEmbeddingBytes []types.DenseEmbeddingByteResult `json:"text_embedding_bytes,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
