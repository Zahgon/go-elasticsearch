package types

type EqlFeaturesJoin struct {
	JoinQueriesFiveOrMore uint `json:"join_queries_five_or_more"`
	JoinQueriesFour       uint `json:"join_queries_four"`
	JoinQueriesThree      uint `json:"join_queries_three"`
	JoinQueriesTwo        uint `json:"join_queries_two"`
	JoinUntil             uint `json:"join_until"`
}

func NewEqlFeaturesJoin() *EqlFeaturesJoin { _ = "STUB: not implemented"; return nil }
