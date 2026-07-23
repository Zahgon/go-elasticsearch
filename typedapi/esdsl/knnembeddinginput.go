package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _knnEmbeddingInput struct {
	v types.KnnEmbeddingInput
}

func NewKnnEmbeddingInput() *_knnEmbeddingInput { _ = "STUB: not implemented"; return nil }

func (u *_knnEmbeddingInput) String(string string) *_knnEmbeddingInput {
	_ = "STUB: not implemented"
	return nil
}

func (u *_knnEmbeddingInput) InferenceStringGroup(inferencestringgroups ...types.InferenceStringVariant) *_knnEmbeddingInput {
	_ = "STUB: not implemented"
	return nil
}

func (u *_knnEmbeddingInput) InferenceStringGroupValues(inferencestringgroupvalues []types.InferenceString) *_knnEmbeddingInput {
	_ = "STUB: not implemented"
	return nil
}

func (u *_inferenceStringGroup) KnnEmbeddingInputCaster() *types.KnnEmbeddingInput {
	_ = "STUB: not implemented"
	return nil
}

func (u *_knnEmbeddingInput) KnnEmbeddingInputCaster() *types.KnnEmbeddingInput {
	_ = "STUB: not implemented"
	return nil
}
