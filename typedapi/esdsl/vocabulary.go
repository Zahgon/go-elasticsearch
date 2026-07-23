package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _vocabulary struct {
	v *types.Vocabulary
}

func NewVocabulary() *_vocabulary { _ = "STUB: not implemented"; return nil }

func (s *_vocabulary) Index(indexname string) *_vocabulary { _ = "STUB: not implemented"; return nil }

func (s *_vocabulary) VocabularyCaster() *types.Vocabulary { _ = "STUB: not implemented"; return nil }
