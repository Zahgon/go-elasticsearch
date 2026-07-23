package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _embeddingStringInput struct {
	v types.EmbeddingStringInput
}

func NewEmbeddingStringInput() *_embeddingStringInput { _ = "STUB: not implemented"; return nil }

func (u *_embeddingStringInput) Strings(strings ...string) *_embeddingStringInput {
	_ = "STUB: not implemented"
	return nil
}

func (u *_embeddingStringInput) EmbeddingStringInputCaster() *types.EmbeddingStringInput {
	_ = "STUB: not implemented"
	return nil
}
