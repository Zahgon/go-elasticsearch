package types

type EmbeddingContentObjectGroup []EmbeddingContentObjectItem

type EmbeddingContentObjectGroupVariant interface {
	EmbeddingContentObjectGroupCaster() *EmbeddingContentObjectGroup
}
