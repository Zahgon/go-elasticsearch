package types

type EmbeddingContentInput []EmbeddingContentObject

type EmbeddingContentInputVariant interface {
	EmbeddingContentInputCaster() *EmbeddingContentInput
}
