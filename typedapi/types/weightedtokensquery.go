package types

type WeightedTokensQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	PruningConfig *TokenPruningConfig `json:"pruning_config,omitempty"`
	QueryName_    *string             `json:"_name,omitempty"`

	Tokens []map[string]float32 `json:"tokens"`
}

func (s *WeightedTokensQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewWeightedTokensQuery() *WeightedTokensQuery { _ = "STUB: not implemented"; return nil }

type WeightedTokensQueryVariant interface {
	WeightedTokensQueryCaster() *WeightedTokensQuery
}

func (s *WeightedTokensQuery) WeightedTokensQueryCaster() *WeightedTokensQuery {
	_ = "STUB: not implemented"
	return nil
}
