package types

type EmbeddingStringInput []string

type EmbeddingStringInputVariant interface {
	EmbeddingStringInputCaster() *EmbeddingStringInput
}
