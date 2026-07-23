package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _property struct {
	v types.Property
}

func NewProperty() *_property { _ = "STUB: not implemented"; return nil }

func (u *_property) UnknownProperty(unknown json.RawMessage) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) BinaryProperty(binaryproperty types.BinaryPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_binaryProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) BooleanProperty(booleanproperty types.BooleanPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_booleanProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) DynamicProperty(dynamicproperty types.DynamicPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_dynamicProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) JoinProperty(joinproperty types.JoinPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_joinProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) KeywordProperty(keywordproperty types.KeywordPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_keywordProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) MatchOnlyTextProperty(matchonlytextproperty types.MatchOnlyTextPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_matchOnlyTextProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) PercolatorProperty(percolatorproperty types.PercolatorPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_percolatorProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) RankFeatureProperty(rankfeatureproperty types.RankFeaturePropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_rankFeatureProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) RankFeaturesProperty(rankfeaturesproperty types.RankFeaturesPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_rankFeaturesProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) SearchAsYouTypeProperty(searchasyoutypeproperty types.SearchAsYouTypePropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_searchAsYouTypeProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) TextProperty(textproperty types.TextPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_textProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) VersionProperty(versionproperty types.VersionPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_versionProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) WildcardProperty(wildcardproperty types.WildcardPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_wildcardProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) DateNanosProperty(datenanosproperty types.DateNanosPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_dateNanosProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) DateProperty(dateproperty types.DatePropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_dateProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) AggregateMetricDoubleProperty(aggregatemetricdoubleproperty types.AggregateMetricDoublePropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_aggregateMetricDoubleProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) DenseVectorProperty(densevectorproperty types.DenseVectorPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_denseVectorProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) FlattenedProperty(flattenedproperty types.FlattenedPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_flattenedProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) NestedProperty(nestedproperty types.NestedPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_nestedProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) ObjectProperty(objectproperty types.ObjectPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_objectProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) PassthroughObjectProperty(passthroughobjectproperty types.PassthroughObjectPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_passthroughObjectProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) RankVectorProperty(rankvectorproperty types.RankVectorPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_rankVectorProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) SemanticTextProperty(semantictextproperty types.SemanticTextPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_semanticTextProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) SparseVectorProperty(sparsevectorproperty types.SparseVectorPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_sparseVectorProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) CompletionProperty(completionproperty types.CompletionPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_completionProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) ConstantKeywordProperty(constantkeywordproperty types.ConstantKeywordPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_constantKeywordProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) CountedKeywordProperty(countedkeywordproperty types.CountedKeywordPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_countedKeywordProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) FieldAliasProperty(fieldaliasproperty types.FieldAliasPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_fieldAliasProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) HistogramProperty(histogramproperty types.HistogramPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_histogramProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) ExponentialHistogramProperty(exponentialhistogramproperty types.ExponentialHistogramPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_exponentialHistogramProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) IpProperty(ipproperty types.IpPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_ipProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) Murmur3HashProperty(murmur3hashproperty types.Murmur3HashPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_murmur3HashProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) TokenCountProperty(tokencountproperty types.TokenCountPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_tokenCountProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) GeoPointProperty(geopointproperty types.GeoPointPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_geoPointProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) GeoShapeProperty(geoshapeproperty types.GeoShapePropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_geoShapeProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) PointProperty(pointproperty types.PointPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_pointProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) ShapeProperty(shapeproperty types.ShapePropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_shapeProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) ByteNumberProperty(bytenumberproperty types.ByteNumberPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_byteNumberProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) DoubleNumberProperty(doublenumberproperty types.DoubleNumberPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_doubleNumberProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) FloatNumberProperty(floatnumberproperty types.FloatNumberPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_floatNumberProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) HalfFloatNumberProperty(halffloatnumberproperty types.HalfFloatNumberPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_halfFloatNumberProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) IntegerNumberProperty(integernumberproperty types.IntegerNumberPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_integerNumberProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) LongNumberProperty(longnumberproperty types.LongNumberPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_longNumberProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) ScaledFloatNumberProperty(scaledfloatnumberproperty types.ScaledFloatNumberPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_scaledFloatNumberProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) ShortNumberProperty(shortnumberproperty types.ShortNumberPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_shortNumberProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) UnsignedLongNumberProperty(unsignedlongnumberproperty types.UnsignedLongNumberPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_unsignedLongNumberProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) DateRangeProperty(daterangeproperty types.DateRangePropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_dateRangeProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) DoubleRangeProperty(doublerangeproperty types.DoubleRangePropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_doubleRangeProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) FloatRangeProperty(floatrangeproperty types.FloatRangePropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_floatRangeProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) IntegerRangeProperty(integerrangeproperty types.IntegerRangePropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_integerRangeProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) IpRangeProperty(iprangeproperty types.IpRangePropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_ipRangeProperty) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }

func (u *_property) LongRangeProperty(longrangeproperty types.LongRangePropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_longRangeProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) IcuCollationProperty(icucollationproperty types.IcuCollationPropertyVariant) *_property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_icuCollationProperty) PropertyCaster() *types.Property {
	_ = "STUB: not implemented"
	return nil
}

func (u *_property) PropertyCaster() *types.Property { _ = "STUB: not implemented"; return nil }
