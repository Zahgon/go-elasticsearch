package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationalternate"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationcasefirst"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationdecomposition"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationstrength"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _icuCollationProperty struct {
	v *types.IcuCollationProperty
}

func NewIcuCollationProperty() *_icuCollationProperty { _ = "STUB: not implemented"; return nil }

func (s *_icuCollationProperty) Alternate(alternate icucollationalternate.IcuCollationAlternate) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) CaseFirst(casefirst icucollationcasefirst.IcuCollationCaseFirst) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) CaseLevel(caselevel bool) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Country(country string) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Decomposition(decomposition icucollationdecomposition.IcuCollationDecomposition) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) HiraganaQuaternaryMode(hiraganaquaternarymode bool) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Index(index bool) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) IndexOptions(indexoptions indexoptions.IndexOptions) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Language(language string) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Norms(norms bool) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) NullValue(nullvalue string) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Numeric(numeric bool) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Rules(rules string) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Strength(strength icucollationstrength.IcuCollationStrength) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) VariableTop(variabletop string) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Variant(variant string) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) CopyTo(fields ...string) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) DocValues(docvalues bool) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Fields(fields map[string]types.Property) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) AddField(key string, value types.PropertyVariant) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) IgnoreAbove(ignoreabove int) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Meta(meta map[string]string) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) AddMeta(key string, value string) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Properties(properties map[string]types.Property) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) AddProperty(key string, value types.PropertyVariant) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) Store(store bool) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_icuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuCollationProperty) IcuCollationPropertyCaster() *types.IcuCollationProperty {
	_ = "STUB: not implemented"
	return nil
}
