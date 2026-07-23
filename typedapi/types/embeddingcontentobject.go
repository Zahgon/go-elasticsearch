package types

type EmbeddingContentObject struct {
	Content []EmbeddingContentObjectItem `json:"content"`
}

func (s *EmbeddingContentObject) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewEmbeddingContentObject() *EmbeddingContentObject { _ = "STUB: not implemented"; return nil }

type EmbeddingContentObjectVariant interface {
	EmbeddingContentObjectCaster() *EmbeddingContentObject
}

func (s *EmbeddingContentObject) EmbeddingContentObjectCaster() *EmbeddingContentObject {
	_ = "STUB: not implemented"
	return nil
}

func (s *EmbeddingContentObject) EmbeddingContentInputCaster() *EmbeddingContentInput {
	_ = "STUB: not implemented"
	return nil
}
