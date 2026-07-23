package types

type EnrichPolicy struct {
	ElasticsearchVersion *string  `json:"elasticsearch_version,omitempty"`
	EnrichFields         []string `json:"enrich_fields"`
	Indices              []string `json:"indices"`
	MatchField           string   `json:"match_field"`
	Name                 *string  `json:"name,omitempty"`
	Query                *Query   `json:"query,omitempty"`
}

func (s *EnrichPolicy) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEnrichPolicy() *EnrichPolicy { _ = "STUB: not implemented"; return nil }

type EnrichPolicyVariant interface {
	EnrichPolicyCaster() *EnrichPolicy
}

func (s *EnrichPolicy) EnrichPolicyCaster() *EnrichPolicy { _ = "STUB: not implemented"; return nil }
