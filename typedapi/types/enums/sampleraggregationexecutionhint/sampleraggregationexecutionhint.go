package sampleraggregationexecutionhint

type SamplerAggregationExecutionHint struct {
	Name string
}

var (
	Map = SamplerAggregationExecutionHint{"map"}

	Globalordinals = SamplerAggregationExecutionHint{"global_ordinals"}

	Byteshash = SamplerAggregationExecutionHint{"bytes_hash"}
)

func (s SamplerAggregationExecutionHint) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SamplerAggregationExecutionHint) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SamplerAggregationExecutionHint) String() string { _ = "STUB: not implemented"; return "" }
