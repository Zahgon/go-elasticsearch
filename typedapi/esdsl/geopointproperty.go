package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geopointmetrictype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _geoPointProperty struct {
	v *types.GeoPointProperty
}

func NewGeoPointProperty() *_geoPointProperty { _ = "STUB: not implemented"; return nil }

func (s *_geoPointProperty) IgnoreMalformed(ignoremalformed bool) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) IgnoreZValue(ignorezvalue bool) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) Index(index bool) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) NullValue(geolocation types.GeoLocationVariant) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) Script(script types.ScriptVariant) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) TimeSeriesMetric(timeseriesmetric geopointmetrictype.GeoPointMetricType) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) CopyTo(fields ...string) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) DocValues(docvalues bool) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) Fields(fields map[string]types.Property) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) AddField(key string, value types.PropertyVariant) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) IgnoreAbove(ignoreabove int) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) Meta(meta map[string]string) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) AddMeta(key string, value string) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) Properties(properties map[string]types.Property) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) AddProperty(key string, value types.PropertyVariant) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) Store(store bool) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_geoPointProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_geoPointProperty) GeoPointPropertyCaster() *types.GeoPointProperty {
	_ = "STUB: not implemented"
	return nil
}
