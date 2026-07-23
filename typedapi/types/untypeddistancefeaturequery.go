package types

import (
	"encoding/json"
)

type UntypedDistanceFeatureQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Field string `json:"field"`

	Origin json.RawMessage `json:"origin,omitempty"`

	Pivot      json.RawMessage `json:"pivot,omitempty"`
	QueryName_ *string         `json:"_name,omitempty"`
}

func (s *UntypedDistanceFeatureQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUntypedDistanceFeatureQuery() *UntypedDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

type UntypedDistanceFeatureQueryVariant interface {
	UntypedDistanceFeatureQueryCaster() *UntypedDistanceFeatureQuery
}

func (s *UntypedDistanceFeatureQuery) UntypedDistanceFeatureQueryCaster() *UntypedDistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *UntypedDistanceFeatureQuery) DistanceFeatureQueryCaster() *DistanceFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}
