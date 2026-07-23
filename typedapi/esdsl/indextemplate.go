package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexTemplate struct {
	v *types.IndexTemplate
}

func NewIndexTemplate() *_indexTemplate { _ = "STUB: not implemented"; return nil }

func (s *_indexTemplate) AllowAutoCreate(allowautocreate bool) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) ComposedOf(composedofs ...string) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) CreatedDate(datetime types.DateTimeVariant) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) CreatedDateMillis(epochtimeunitmillis int64) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) DataStream(datastream types.IndexTemplateDataStreamConfigurationVariant) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) Deprecated(deprecated bool) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) IgnoreMissingComponentTemplates(names ...string) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) IndexPatterns(names ...string) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) Meta_(metadata types.MetadataVariant) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) ModifiedDate(datetime types.DateTimeVariant) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) ModifiedDateMillis(epochtimeunitmillis int64) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) Priority(priority int64) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) Template(template types.IndexTemplateSummaryVariant) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) Version(versionnumber int64) *_indexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexTemplate) IndexTemplateCaster() *types.IndexTemplate {
	_ = "STUB: not implemented"
	return nil
}
