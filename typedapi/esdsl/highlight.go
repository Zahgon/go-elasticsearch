package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/boundaryscanner"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlighterencoder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlighterfragmenter"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlighterorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlightertagsschema"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlightertype"
)

type _highlight struct {
	v *types.Highlight
}

func NewHighlight() *_highlight { _ = "STUB: not implemented"; return nil }

func (s *_highlight) Encoder(encoder highlighterencoder.HighlighterEncoder) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) Fields(fields []map[string]types.HighlightField) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) BoundaryChars(boundarychars string) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) BoundaryMaxScan(boundarymaxscan int) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) BoundaryScanner(boundaryscanner boundaryscanner.BoundaryScanner) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) BoundaryScannerLocale(boundaryscannerlocale string) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) ForceSource(forcesource bool) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) FragmentSize(fragmentsize int) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) Fragmenter(fragmenter highlighterfragmenter.HighlighterFragmenter) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) HighlightFilter(highlightfilter bool) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) HighlightQuery(highlightquery types.QueryVariant) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) MaxAnalyzedOffset(maxanalyzedoffset int) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) MaxFragmentLength(maxfragmentlength int) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) NoMatchSize(nomatchsize int) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) NumberOfFragments(numberoffragments int) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) Options(options map[string]json.RawMessage) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) AddOption(key string, value json.RawMessage) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) Order(order highlighterorder.HighlighterOrder) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) PhraseLimit(phraselimit int) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) PostTags(posttags ...string) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) PreTags(pretags ...string) *_highlight { _ = "STUB: not implemented"; return nil }

func (s *_highlight) RequireFieldMatch(requirefieldmatch bool) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) TagsSchema(tagsschema highlightertagsschema.HighlighterTagsSchema) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) Type(type_ highlightertype.HighlighterType) *_highlight {
	_ = "STUB: not implemented"
	return nil
}

func (s *_highlight) HighlightCaster() *types.Highlight { _ = "STUB: not implemented"; return nil }
