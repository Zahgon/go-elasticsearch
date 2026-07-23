package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termvectoroption"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeseriesmetrictype"
)

type _dynamicProperty struct {
	v *types.DynamicProperty
}

func NewDynamicProperty() *_dynamicProperty { _ = "STUB: not implemented"; return nil }

func (s *_dynamicProperty) Analyzer(analyzer string) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) Boost(boost types.Float64) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) Coerce(coerce bool) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) EagerGlobalOrdinals(eagerglobalordinals bool) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) Enabled(enabled bool) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) Format(format string) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) IgnoreMalformed(ignoremalformed bool) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) Index(index bool) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) IndexOptions(indexoptions indexoptions.IndexOptions) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) IndexPhrases(indexphrases bool) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) IndexPrefixes(indexprefixes types.TextIndexPrefixesVariant) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) Locale(locale string) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) Norms(norms bool) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) NullValue(fieldvalue types.FieldValueVariant) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) PositionIncrementGap(positionincrementgap int) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) PrecisionStep(precisionstep int) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) Script(script types.ScriptVariant) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) SearchAnalyzer(searchanalyzer string) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) SearchQuoteAnalyzer(searchquoteanalyzer string) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) TermVector(termvector termvectoroption.TermVectorOption) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) CopyTo(fields ...string) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) DocValues(docvalues bool) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) Fields(fields map[string]types.Property) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) AddField(key string, value types.PropertyVariant) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) IgnoreAbove(ignoreabove int) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) Meta(meta map[string]string) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) AddMeta(key string, value string) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) Properties(properties map[string]types.Property) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) AddProperty(key string, value types.PropertyVariant) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) Store(store bool) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_dynamicProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicProperty) DynamicPropertyCaster() *types.DynamicProperty {
	_ = "STUB: not implemented"
	return nil
}
