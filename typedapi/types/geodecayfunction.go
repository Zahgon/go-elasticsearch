package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

type GeoDecayFunction struct {
	DecayFunctionBaseGeoLocationDistance map[string]DecayPlacementGeoLocationDistance `json:"-"`

	MultiValueMode *multivaluemode.MultiValueMode `json:"multi_value_mode,omitempty"`
}

func (s *GeoDecayFunction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GeoDecayFunction) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGeoDecayFunction() *GeoDecayFunction { _ = "STUB: not implemented"; return nil }

type GeoDecayFunctionVariant interface {
	GeoDecayFunctionCaster() *GeoDecayFunction
}

func (s *GeoDecayFunction) GeoDecayFunctionCaster() *GeoDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *GeoDecayFunction) DecayFunctionCaster() *DecayFunction {
	_ = "STUB: not implemented"
	return nil
}
