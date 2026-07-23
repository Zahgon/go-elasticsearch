package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
)

type _commonTermsQuery struct {
	k string
	v *types.CommonTermsQuery
}

func NewCommonTermsQuery(field string, query string) *_commonTermsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonTermsQuery) Analyzer(analyzer string) *_commonTermsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonTermsQuery) CutoffFrequency(cutofffrequency types.Float64) *_commonTermsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonTermsQuery) HighFreqOperator(highfreqoperator operator.Operator) *_commonTermsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonTermsQuery) LowFreqOperator(lowfreqoperator operator.Operator) *_commonTermsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonTermsQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_commonTermsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonTermsQuery) Query(query string) *_commonTermsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonTermsQuery) Boost(boost float32) *_commonTermsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonTermsQuery) QueryName_(queryname_ string) *_commonTermsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonTermsQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func NewSingleCommonTermsQuery() *_commonTermsQuery { _ = "STUB: not implemented"; return nil }

func (s *_commonTermsQuery) CommonTermsQueryCaster() *types.CommonTermsQuery {
	_ = "STUB: not implemented"
	return nil
}
