package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortmode"
)

type MatrixStatsAggregation struct {
	Fields []string `json:"fields,omitempty"`

	Missing map[string]Float64 `json:"missing,omitempty"`

	Mode *sortmode.SortMode `json:"mode,omitempty"`
}

func (s *MatrixStatsAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMatrixStatsAggregation() *MatrixStatsAggregation { _ = "STUB: not implemented"; return nil }

type MatrixStatsAggregationVariant interface {
	MatrixStatsAggregationCaster() *MatrixStatsAggregation
}

func (s *MatrixStatsAggregation) MatrixStatsAggregationCaster() *MatrixStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}
