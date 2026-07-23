package types

type HtmlStripCharFilter struct {
	EscapedTags []string `json:"escaped_tags,omitempty"`
	Type        string   `json:"type,omitempty"`
	Version     *string  `json:"version,omitempty"`
}

func (s *HtmlStripCharFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s HtmlStripCharFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewHtmlStripCharFilter() *HtmlStripCharFilter { _ = "STUB: not implemented"; return nil }

type HtmlStripCharFilterVariant interface {
	HtmlStripCharFilterCaster() *HtmlStripCharFilter
}

func (s *HtmlStripCharFilter) HtmlStripCharFilterCaster() *HtmlStripCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *HtmlStripCharFilter) CharFilterDefinitionCaster() *CharFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
