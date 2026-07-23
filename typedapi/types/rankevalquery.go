package types

type RankEvalQuery struct {
	Query Query `json:"query"`
	Size  *int  `json:"size,omitempty"`
}

func (s *RankEvalQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRankEvalQuery() *RankEvalQuery { _ = "STUB: not implemented"; return nil }

type RankEvalQueryVariant interface {
	RankEvalQueryCaster() *RankEvalQuery
}

func (s *RankEvalQuery) RankEvalQueryCaster() *RankEvalQuery { _ = "STUB: not implemented"; return nil }
