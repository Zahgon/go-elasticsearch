package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _enrichPolicy struct {
	v *types.EnrichPolicy
}

func NewEnrichPolicy() *_enrichPolicy { _ = "STUB: not implemented"; return nil }

func (s *_enrichPolicy) ElasticsearchVersion(elasticsearchversion string) *_enrichPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichPolicy) EnrichFields(fields ...string) *_enrichPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichPolicy) Indices(indices ...string) *_enrichPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichPolicy) MatchField(field string) *_enrichPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichPolicy) Name(name string) *_enrichPolicy { _ = "STUB: not implemented"; return nil }

func (s *_enrichPolicy) Query(query types.QueryVariant) *_enrichPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_enrichPolicy) EnrichPolicyCaster() *types.EnrichPolicy {
	_ = "STUB: not implemented"
	return nil
}
