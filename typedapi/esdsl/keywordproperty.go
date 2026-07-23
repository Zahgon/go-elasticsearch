package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _keywordProperty struct {
	v *types.KeywordProperty
}

func NewKeywordProperty() *_keywordProperty { _ = "STUB: not implemented"; return nil }

func (s *_keywordProperty) Boost(boost types.Float64) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) EagerGlobalOrdinals(eagerglobalordinals bool) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) Index(index bool) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) IndexOptions(indexoptions indexoptions.IndexOptions) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) Normalizer(normalizer string) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) Norms(norms bool) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) NullValue(nullvalue string) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) OnScriptError(onscripterror onscripterror.OnScriptError) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) Script(script types.ScriptVariant) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) Similarity(similarity string) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) SplitQueriesOnWhitespace(splitqueriesonwhitespace bool) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) TimeSeriesDimension(timeseriesdimension bool) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) CopyTo(fields ...string) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) DocValues(docvalues bool) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) Fields(fields map[string]types.Property) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) AddField(key string, value types.PropertyVariant) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) IgnoreAbove(ignoreabove int) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) Meta(meta map[string]string) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) AddMeta(key string, value string) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) Properties(properties map[string]types.Property) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) AddProperty(key string, value types.PropertyVariant) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) Store(store bool) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_keywordProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordProperty) KeywordPropertyCaster() *types.KeywordProperty {
	_ = "STUB: not implemented"
	return nil
}
