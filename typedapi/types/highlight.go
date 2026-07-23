package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/boundaryscanner"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlighterencoder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlighterfragmenter"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlighterorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlightertagsschema"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlightertype"
)

type Highlight struct {
	BoundaryChars *string `json:"boundary_chars,omitempty"`

	BoundaryMaxScan *int `json:"boundary_max_scan,omitempty"`

	BoundaryScanner *boundaryscanner.BoundaryScanner `json:"boundary_scanner,omitempty"`

	BoundaryScannerLocale *string                                `json:"boundary_scanner_locale,omitempty"`
	Encoder               *highlighterencoder.HighlighterEncoder `json:"encoder,omitempty"`
	Fields                []map[string]HighlightField            `json:"fields"`
	ForceSource           *bool                                  `json:"force_source,omitempty"`

	FragmentSize *int `json:"fragment_size,omitempty"`

	Fragmenter      *highlighterfragmenter.HighlighterFragmenter `json:"fragmenter,omitempty"`
	HighlightFilter *bool                                        `json:"highlight_filter,omitempty"`

	HighlightQuery *Query `json:"highlight_query,omitempty"`

	MaxAnalyzedOffset *int `json:"max_analyzed_offset,omitempty"`
	MaxFragmentLength *int `json:"max_fragment_length,omitempty"`

	NoMatchSize *int `json:"no_match_size,omitempty"`

	NumberOfFragments *int                       `json:"number_of_fragments,omitempty"`
	Options           map[string]json.RawMessage `json:"options,omitempty"`

	Order *highlighterorder.HighlighterOrder `json:"order,omitempty"`

	PhraseLimit *int `json:"phrase_limit,omitempty"`

	PostTags []string `json:"post_tags,omitempty"`

	PreTags []string `json:"pre_tags,omitempty"`

	RequireFieldMatch *bool `json:"require_field_match,omitempty"`

	TagsSchema *highlightertagsschema.HighlighterTagsSchema `json:"tags_schema,omitempty"`
	Type       *highlightertype.HighlighterType             `json:"type,omitempty"`
}

func (s *Highlight) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHighlight() *Highlight { _ = "STUB: not implemented"; return nil }

type HighlightVariant interface {
	HighlightCaster() *Highlight
}

func (s *Highlight) HighlightCaster() *Highlight { _ = "STUB: not implemented"; return nil }
