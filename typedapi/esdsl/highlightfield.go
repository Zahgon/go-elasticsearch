package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/boundaryscanner"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlighterfragmenter"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlighterorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlightertagsschema"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlightertype"
)

type _highlightField struct {
	v *types.HighlightField
}

func NewHighlightField() *_highlightField { _ = "STUB: not implemented"; return nil }

func (s *_highlightField) FragmentOffset(fragmentoffset int) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) MatchedFields(fields ...string) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) BoundaryChars(boundarychars string) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) BoundaryMaxScan(boundarymaxscan int) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) BoundaryScanner(boundaryscanner boundaryscanner.BoundaryScanner) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) BoundaryScannerLocale(boundaryscannerlocale string) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) ForceSource(forcesource bool) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) FragmentSize(fragmentsize int) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) Fragmenter(fragmenter highlighterfragmenter.HighlighterFragmenter) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) HighlightFilter(highlightfilter bool) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) HighlightQuery(highlightquery types.QueryVariant) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) MaxAnalyzedOffset(maxanalyzedoffset int) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) MaxFragmentLength(maxfragmentlength int) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) NoMatchSize(nomatchsize int) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) NumberOfFragments(numberoffragments int) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) Options(options map[string]json.RawMessage) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) AddOption(key string, value json.RawMessage) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) Order(order highlighterorder.HighlighterOrder) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) PhraseLimit(phraselimit int) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) PostTags(posttags ...string) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) PreTags(pretags ...string) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) RequireFieldMatch(requirefieldmatch bool) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) TagsSchema(tagsschema highlightertagsschema.HighlighterTagsSchema) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) Type(type_ highlightertype.HighlighterType) *_highlightField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlightField) HighlightFieldCaster() *types.HighlightField {
	_ = "STUB: not implemented"
	return nil
}
