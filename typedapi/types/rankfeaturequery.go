package types

type RankFeatureQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Field string `json:"field"`

	Linear *RankFeatureFunctionLinear `json:"linear,omitempty"`

	Log        *RankFeatureFunctionLogarithm `json:"log,omitempty"`
	QueryName_ *string                       `json:"_name,omitempty"`

	Saturation *RankFeatureFunctionSaturation `json:"saturation,omitempty"`

	Sigmoid *RankFeatureFunctionSigmoid `json:"sigmoid,omitempty"`
}

func (s *RankFeatureQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRankFeatureQuery() *RankFeatureQuery { _ = "STUB: not implemented"; return nil }

type RankFeatureQueryVariant interface {
	RankFeatureQueryCaster() *RankFeatureQuery
}

func (s *RankFeatureQuery) RankFeatureQueryCaster() *RankFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}
