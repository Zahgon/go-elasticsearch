package types

type EqlFeaturesSequences struct {
	SequenceMaxspan           uint `json:"sequence_maxspan"`
	SequenceQueriesFiveOrMore uint `json:"sequence_queries_five_or_more"`
	SequenceQueriesFour       uint `json:"sequence_queries_four"`
	SequenceQueriesThree      uint `json:"sequence_queries_three"`
	SequenceQueriesTwo        uint `json:"sequence_queries_two"`
	SequenceUntil             uint `json:"sequence_until"`
}

func NewEqlFeaturesSequences() *EqlFeaturesSequences { _ = "STUB: not implemented"; return nil }
