package types

type KnnEmbeddingInput any

type KnnEmbeddingInputVariant interface {
	KnnEmbeddingInputCaster() *KnnEmbeddingInput
}
