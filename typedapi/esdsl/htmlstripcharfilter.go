package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _htmlStripCharFilter struct {
	v *types.HtmlStripCharFilter
}

func NewHtmlStripCharFilter() *_htmlStripCharFilter { _ = "STUB: not implemented"; return nil }

func (s *_htmlStripCharFilter) EscapedTags(escapedtags ...string) *_htmlStripCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_htmlStripCharFilter) Version(versionstring string) *_htmlStripCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_htmlStripCharFilter) HtmlStripCharFilterCaster() *types.HtmlStripCharFilter {
	_ = "STUB: not implemented"
	return nil
}
