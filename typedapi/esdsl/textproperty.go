package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termvectoroption"
)

type _textProperty struct {
	v *types.TextProperty
}

func NewTextProperty() *_textProperty { _ = "STUB: not implemented"; return nil }

func (s *_textProperty) Analyzer(analyzer string) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) Boost(boost types.Float64) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) EagerGlobalOrdinals(eagerglobalordinals bool) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) Fielddata(fielddata bool) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) FielddataFrequencyFilter(fielddatafrequencyfilter types.FielddataFrequencyFilterVariant) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) Index(index bool) *_textProperty { _ = "STUB: not implemented"; return nil }

func (s *_textProperty) IndexOptions(indexoptions indexoptions.IndexOptions) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) IndexPhrases(indexphrases bool) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) IndexPrefixes(indexprefixes types.TextIndexPrefixesVariant) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) Norms(norms bool) *_textProperty { _ = "STUB: not implemented"; return nil }

func (s *_textProperty) PositionIncrementGap(positionincrementgap int) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) SearchAnalyzer(searchanalyzer string) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) SearchQuoteAnalyzer(searchquoteanalyzer string) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) Similarity(similarity string) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) TermVector(termvector termvectoroption.TermVectorOption) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) CopyTo(fields ...string) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) Fields(fields map[string]types.Property) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) AddField(key string, value types.PropertyVariant) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) IgnoreAbove(ignoreabove int) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) Meta(meta map[string]string) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) AddMeta(key string, value string) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) Properties(properties map[string]types.Property) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) AddProperty(key string, value types.PropertyVariant) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) Store(store bool) *_textProperty { _ = "STUB: not implemented"; return nil }

func (s *_textProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_textProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textProperty) TextPropertyCaster() *types.TextProperty {
	_ = "STUB: not implemented"
	return nil
}
