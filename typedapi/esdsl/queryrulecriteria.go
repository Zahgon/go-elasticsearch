package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/queryrulecriteriatype"
)

type _queryRuleCriteria struct {
	v *types.QueryRuleCriteria
}

func NewQueryRuleCriteria(type_ queryrulecriteriatype.QueryRuleCriteriaType) *_queryRuleCriteria {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryRuleCriteria) Metadata(metadata string) *_queryRuleCriteria {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryRuleCriteria) Type(type_ queryrulecriteriatype.QueryRuleCriteriaType) *_queryRuleCriteria {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryRuleCriteria) Values(values ...json.RawMessage) *_queryRuleCriteria {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryRuleCriteria) QueryRuleCriteriaCaster() *types.QueryRuleCriteria {
	_ = "STUB: not implemented"
	return nil
}
