package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationexecutionhint"
)

type SignificantTermsAggregation struct {
	BackgroundFilter *Query `json:"background_filter,omitempty"`

	ChiSquare *ChiSquareHeuristic `json:"chi_square,omitempty"`

	Exclude []string `json:"exclude,omitempty"`

	ExecutionHint *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`

	Field *string `json:"field,omitempty"`

	Gnd *GoogleNormalizedDistanceHeuristic `json:"gnd,omitempty"`

	Include TermsInclude `json:"include,omitempty"`

	Jlh *EmptyObject `json:"jlh,omitempty"`

	MinDocCount *int64 `json:"min_doc_count,omitempty"`

	MutualInformation *MutualInformationHeuristic `json:"mutual_information,omitempty"`

	PValue *PValueHeuristic `json:"p_value,omitempty"`

	Percentage *PercentageScoreHeuristic `json:"percentage,omitempty"`

	ScriptHeuristic *ScriptedHeuristic `json:"script_heuristic,omitempty"`

	ShardMinDocCount *int64 `json:"shard_min_doc_count,omitempty"`

	ShardSize *int `json:"shard_size,omitempty"`

	Size *int `json:"size,omitempty"`
}

func (s *SignificantTermsAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSignificantTermsAggregation() *SignificantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

type SignificantTermsAggregationVariant interface {
	SignificantTermsAggregationCaster() *SignificantTermsAggregation
}

func (s *SignificantTermsAggregation) SignificantTermsAggregationCaster() *SignificantTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}
