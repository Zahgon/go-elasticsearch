package types

type TopHitsAggregation struct {
	DocvalueFields []FieldAndFormat `json:"docvalue_fields,omitempty"`

	Explain *bool `json:"explain,omitempty"`

	Field *string `json:"field,omitempty"`

	Fields []FieldAndFormat `json:"fields,omitempty"`

	From *int `json:"from,omitempty"`

	Highlight *Highlight `json:"highlight,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`

	ScriptFields map[string]ScriptField `json:"script_fields,omitempty"`

	SeqNoPrimaryTerm *bool `json:"seq_no_primary_term,omitempty"`

	Size *int `json:"size,omitempty"`

	Sort []SortCombinations `json:"sort,omitempty"`

	Source_ SourceConfig `json:"_source,omitempty"`

	StoredFields []string `json:"stored_fields,omitempty"`

	TrackScores *bool `json:"track_scores,omitempty"`

	Version *bool `json:"version,omitempty"`
}

func (s *TopHitsAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTopHitsAggregation() *TopHitsAggregation { _ = "STUB: not implemented"; return nil }

type TopHitsAggregationVariant interface {
	TopHitsAggregationCaster() *TopHitsAggregation
}

func (s *TopHitsAggregation) TopHitsAggregationCaster() *TopHitsAggregation {
	_ = "STUB: not implemented"
	return nil
}
