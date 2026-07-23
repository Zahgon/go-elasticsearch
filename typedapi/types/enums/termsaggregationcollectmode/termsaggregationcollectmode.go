package termsaggregationcollectmode

type TermsAggregationCollectMode struct {
	Name string
}

var (
	Depthfirst = TermsAggregationCollectMode{"depth_first"}

	Breadthfirst = TermsAggregationCollectMode{"breadth_first"}
)

func (t TermsAggregationCollectMode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TermsAggregationCollectMode) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TermsAggregationCollectMode) String() string { _ = "STUB: not implemented"; return "" }
