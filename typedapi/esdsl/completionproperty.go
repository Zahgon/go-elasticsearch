package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
)

type _completionProperty struct {
	v *types.CompletionProperty
}

func NewCompletionProperty() *_completionProperty { _ = "STUB: not implemented"; return nil }

func (s *_completionProperty) Analyzer(analyzer string) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) Contexts(contexts ...types.SuggestContextVariant) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) ContextsValues(contextsvalues []types.SuggestContext) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) MaxInputLength(maxinputlength int) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) PreservePositionIncrements(preservepositionincrements bool) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) PreserveSeparators(preserveseparators bool) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) SearchAnalyzer(searchanalyzer string) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) CopyTo(fields ...string) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) DocValues(docvalues bool) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) Fields(fields map[string]types.Property) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) AddField(key string, value types.PropertyVariant) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) IgnoreAbove(ignoreabove int) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) Meta(meta map[string]string) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) AddMeta(key string, value string) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) Properties(properties map[string]types.Property) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) AddProperty(key string, value types.PropertyVariant) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) Store(store bool) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_completionProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionProperty) CompletionPropertyCaster() *types.CompletionProperty {
	_ = "STUB: not implemented"
	return nil
}
