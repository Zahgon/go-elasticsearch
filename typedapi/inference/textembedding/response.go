package textembedding

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	TextEmbedding      []types.DenseEmbeddingResult     `json:"text_embedding,omitempty"`
	TextEmbeddingBits  []types.DenseEmbeddingByteResult `json:"text_embedding_bits,omitempty"`
	TextEmbeddingBytes []types.DenseEmbeddingByteResult `json:"text_embedding_bytes,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
