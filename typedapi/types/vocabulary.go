package types

type Vocabulary struct {
	Index string `json:"index"`
}

func (s *Vocabulary) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewVocabulary() *Vocabulary { _ = "STUB: not implemented"; return nil }

type VocabularyVariant interface {
	VocabularyCaster() *Vocabulary
}

func (s *Vocabulary) VocabularyCaster() *Vocabulary { _ = "STUB: not implemented"; return nil }
