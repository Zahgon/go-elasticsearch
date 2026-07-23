package types

type EmbeddingInput any

type EmbeddingInputVariant interface {
	EmbeddingInputCaster() *EmbeddingInput
}
