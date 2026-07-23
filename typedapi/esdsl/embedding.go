package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _embedding struct {
	v *types.Embedding
}

func NewEmbedding() *_embedding { _ = "STUB: not implemented"; return nil }

func (s *_embedding) InferenceId(inferenceid string) *_embedding {
	_ = "STUB: not implemented"
	return nil
}

func (s *_embedding) Input(knnembeddinginput types.KnnEmbeddingInputVariant) *_embedding {
	_ = "STUB: not implemented"
	return nil
}

func (s *_embedding) Timeout(duration types.DurationVariant) *_embedding {
	_ = "STUB: not implemented"
	return nil
}

func (s *_embedding) QueryVectorBuilderCaster() *types.QueryVectorBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *_embedding) EmbeddingCaster() *types.Embedding { _ = "STUB: not implemented"; return nil }
