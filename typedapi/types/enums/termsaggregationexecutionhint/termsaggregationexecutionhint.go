package termsaggregationexecutionhint

type TermsAggregationExecutionHint struct {
	Name string
}

var (
	Map = TermsAggregationExecutionHint{"map"}

	Globalordinals = TermsAggregationExecutionHint{"global_ordinals"}

	Globalordinalshash = TermsAggregationExecutionHint{"global_ordinals_hash"}

	Globalordinalslowcardinality = TermsAggregationExecutionHint{"global_ordinals_low_cardinality"}
)

func (t TermsAggregationExecutionHint) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TermsAggregationExecutionHint) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TermsAggregationExecutionHint) String() string { _ = "STUB: not implemented"; return "" }
