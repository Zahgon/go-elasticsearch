package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _embeddingInput struct {
	v types.EmbeddingInput
}

func NewEmbeddingInput() *_embeddingInput { _ = "STUB: not implemented"; return nil }

func (u *_embeddingInput) EmbeddingStringInput(embeddingstringinputs ...string) *_embeddingInput {
	_ = "STUB: not implemented"
	return nil
}

func (u *_embeddingStringInput) EmbeddingInputCaster() *types.EmbeddingInput {
	_ = "STUB: not implemented"
	return nil
}

func (u *_embeddingInput) EmbeddingContentInput(embeddingcontentinputs ...types.EmbeddingContentObjectVariant) *_embeddingInput {
	_ = "STUB: not implemented"
	return nil
}

func (u *_embeddingInput) EmbeddingContentInputValues(embeddingcontentinputvalues []types.EmbeddingContentObject) *_embeddingInput {
	_ = "STUB: not implemented"
	return nil
}

func (u *_embeddingContentInput) EmbeddingInputCaster() *types.EmbeddingInput {
	_ = "STUB: not implemented"
	return nil
}

func (u *_embeddingInput) EmbeddingInputCaster() *types.EmbeddingInput {
	_ = "STUB: not implemented"
	return nil
}
