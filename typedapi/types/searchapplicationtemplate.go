package types

type SearchApplicationTemplate struct {
	Script Script `json:"script"`
}

func NewSearchApplicationTemplate() *SearchApplicationTemplate {
	_ = "STUB: not implemented"
	return nil
}

type SearchApplicationTemplateVariant interface {
	SearchApplicationTemplateCaster() *SearchApplicationTemplate
}

func (s *SearchApplicationTemplate) SearchApplicationTemplateCaster() *SearchApplicationTemplate {
	_ = "STUB: not implemented"
	return nil
}
