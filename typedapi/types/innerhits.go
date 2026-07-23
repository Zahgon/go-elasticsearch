package types

type InnerHits struct {
	Collapse       *FieldCollapse   `json:"collapse,omitempty"`
	DocvalueFields []FieldAndFormat `json:"docvalue_fields,omitempty"`
	Explain        *bool            `json:"explain,omitempty"`
	Field          []string         `json:"field,omitempty"`
	Fields         []FieldAndFormat `json:"fields,omitempty"`

	From           *int       `json:"from,omitempty"`
	Highlight      *Highlight `json:"highlight,omitempty"`
	IgnoreUnmapped *bool      `json:"ignore_unmapped,omitempty"`

	Name             *string                `json:"name,omitempty"`
	ScriptFields     map[string]ScriptField `json:"script_fields,omitempty"`
	SeqNoPrimaryTerm *bool                  `json:"seq_no_primary_term,omitempty"`

	Size *int `json:"size,omitempty"`

	Sort         []SortCombinations `json:"sort,omitempty"`
	Source_      SourceConfig       `json:"_source,omitempty"`
	StoredFields []string           `json:"stored_fields,omitempty"`
	TrackScores  *bool              `json:"track_scores,omitempty"`
	Version      *bool              `json:"version,omitempty"`
}

func (s *InnerHits) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewInnerHits() *InnerHits { _ = "STUB: not implemented"; return nil }

type InnerHitsVariant interface {
	InnerHitsCaster() *InnerHits
}

func (s *InnerHits) InnerHitsCaster() *InnerHits { _ = "STUB: not implemented"; return nil }
