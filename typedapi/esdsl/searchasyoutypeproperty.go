package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syntheticsourcekeepenum"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termvectoroption"
)

type _searchAsYouTypeProperty struct {
	v *types.SearchAsYouTypeProperty
}

func NewSearchAsYouTypeProperty() *_searchAsYouTypeProperty { _ = "STUB: not implemented"; return nil }

func (s *_searchAsYouTypeProperty) Analyzer(analyzer string) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) Index(index bool) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) IndexOptions(indexoptions indexoptions.IndexOptions) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) MaxShingleSize(maxshinglesize int) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) Norms(norms bool) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) SearchAnalyzer(searchanalyzer string) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) SearchQuoteAnalyzer(searchquoteanalyzer string) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) Similarity(similarity string) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) TermVector(termvector termvectoroption.TermVectorOption) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) CopyTo(fields ...string) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) Dynamic(dynamic dynamicmapping.DynamicMapping) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) Fields(fields map[string]types.Property) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) AddField(key string, value types.PropertyVariant) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) IgnoreAbove(ignoreabove int) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) Meta(meta map[string]string) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) AddMeta(key string, value string) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) Properties(properties map[string]types.Property) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) AddProperty(key string, value types.PropertyVariant) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) Store(store bool) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) SyntheticSourceKeep(syntheticsourcekeep syntheticsourcekeepenum.SyntheticSourceKeepEnum) *_searchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchAsYouTypeProperty) SearchAsYouTypePropertyCaster() *types.SearchAsYouTypeProperty {
	_ = "STUB: not implemented"
	return nil
}
