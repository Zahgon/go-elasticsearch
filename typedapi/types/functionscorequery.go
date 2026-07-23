package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/functionboostmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/functionscoremode"
)

type FunctionScoreQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	BoostMode *functionboostmode.FunctionBoostMode `json:"boost_mode,omitempty"`

	Functions []FunctionScore `json:"functions,omitempty"`

	MaxBoost *Float64 `json:"max_boost,omitempty"`

	MinScore *Float64 `json:"min_score,omitempty"`

	Query      *Query  `json:"query,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`

	ScoreMode *functionscoremode.FunctionScoreMode `json:"score_mode,omitempty"`
}

func (s *FunctionScoreQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFunctionScoreQuery() *FunctionScoreQuery { _ = "STUB: not implemented"; return nil }

type FunctionScoreQueryVariant interface {
	FunctionScoreQueryCaster() *FunctionScoreQuery
}

func (s *FunctionScoreQuery) FunctionScoreQueryCaster() *FunctionScoreQuery {
	_ = "STUB: not implemented"
	return nil
}
