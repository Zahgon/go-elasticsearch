package shardsstatsstage

type ShardsStatsStage struct {
	Name string
}

var (
	DONE = ShardsStatsStage{"DONE"}

	FAILURE = ShardsStatsStage{"FAILURE"}

	FINALIZE = ShardsStatsStage{"FINALIZE"}

	INIT = ShardsStatsStage{"INIT"}

	STARTED = ShardsStatsStage{"STARTED"}
)

func (s ShardsStatsStage) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ShardsStatsStage) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s ShardsStatsStage) String() string { _ = "STUB: not implemented"; return "" }
