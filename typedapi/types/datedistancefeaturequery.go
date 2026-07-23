package types

type DateDistanceFeatureQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Field string `json:"field"`

	Origin string `json:"origin"`

	Pivot      Duration `json:"pivot"`
	QueryName_ *string  `json:"_name,omitempty"`
}

func (s *DateDistanceFeatureQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDateDistanceFeatureQuery() *DateDistanceFeatureQuery { _ = "STUB: not implemented"; return nil }

type DateDistanceFeatureQueryVariant interface {
	DateDistanceFeatureQueryCaster() *DateDistanceFeatureQuery
}

func (s *DateDistanceFeatureQuery) DateDistanceFeatureQueryCaster() *DateDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *DateDistanceFeatureQuery) DistanceFeatureQueryCaster() *DistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}
