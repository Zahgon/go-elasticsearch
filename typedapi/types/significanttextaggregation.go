package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationexecutionhint"
)

type SignificantTextAggregation struct {
	BackgroundFilter *Query `json:"background_filter,omitempty"`

	ChiSquare *ChiSquareHeuristic `json:"chi_square,omitempty"`

	Exclude []string `json:"exclude,omitempty"`

	ExecutionHint *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`

	Field *string `json:"field,omitempty"`

	FilterDuplicateText *bool `json:"filter_duplicate_text,omitempty"`

	Gnd *GoogleNormalizedDistanceHeuristic `json:"gnd,omitempty"`

	Include TermsInclude `json:"include,omitempty"`

	Jlh *EmptyObject `json:"jlh,omitempty"`

	MinDocCount *int64 `json:"min_doc_count,omitempty"`

	MutualInformation *MutualInformationHeuristic `json:"mutual_information,omitempty"`

	Percentage *PercentageScoreHeuristic `json:"percentage,omitempty"`

	ScriptHeuristic *ScriptedHeuristic `json:"script_heuristic,omitempty"`

	ShardMinDocCount *int64 `json:"shard_min_doc_count,omitempty"`

	ShardSize *int `json:"shard_size,omitempty"`

	Size *int `json:"size,omitempty"`

	SourceFields []string `json:"source_fields,omitempty"`
}

func (s *SignificantTextAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSignificantTextAggregation() *SignificantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}

type SignificantTextAggregationVariant interface {
	SignificantTextAggregationCaster() *SignificantTextAggregation
}

func (s *SignificantTextAggregation) SignificantTextAggregationCaster() *SignificantTextAggregation {
	_ = "STUB: not implemented"
	return nil
}
