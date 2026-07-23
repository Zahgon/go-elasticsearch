package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _phraseSuggestCollate struct {
	v *types.PhraseSuggestCollate
}

func NewPhraseSuggestCollate(query types.PhraseSuggestCollateQueryVariant) *_phraseSuggestCollate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggestCollate) Params(params map[string]json.RawMessage) *_phraseSuggestCollate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggestCollate) AddParam(key string, value json.RawMessage) *_phraseSuggestCollate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggestCollate) Prune(prune bool) *_phraseSuggestCollate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggestCollate) Query(query types.PhraseSuggestCollateQueryVariant) *_phraseSuggestCollate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggestCollate) PhraseSuggestCollateCaster() *types.PhraseSuggestCollate {
	_ = "STUB: not implemented"
	return nil
}
