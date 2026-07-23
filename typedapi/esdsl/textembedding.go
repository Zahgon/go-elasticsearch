package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textEmbedding struct {
	v *types.TextEmbedding
}

func NewTextEmbedding(modeltext string) *_textEmbedding { _ = "STUB: not implemented"; return nil }

func (s *_textEmbedding) ModelId(modelid string) *_textEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textEmbedding) ModelText(modeltext string) *_textEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textEmbedding) QueryVectorBuilderCaster() *types.QueryVectorBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textEmbedding) TextEmbeddingCaster() *types.TextEmbedding {
	_ = "STUB: not implemented"
	return nil
}
