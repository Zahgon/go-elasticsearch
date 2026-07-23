package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

type _geoDecayFunction struct {
	v *types.GeoDecayFunction
}

func NewGeoDecayFunction() *_geoDecayFunction { _ = "STUB: not implemented"; return nil }

func (s *_geoDecayFunction) DecayFunctionBaseGeoLocationDistance(decayfunctionbasegeolocationdistance map[string]types.DecayPlacementGeoLocationDistance) *_geoDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDecayFunction) AddDecayFunctionBaseGeoLocationDistance(key string, value types.DecayPlacementGeoLocationDistanceVariant) *_geoDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDecayFunction) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *_geoDecayFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDecayFunction) FunctionScoreCaster() *types.FunctionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoDecayFunction) GeoDecayFunctionCaster() *types.GeoDecayFunction {
	_ = "STUB: not implemented"
	return nil
}
